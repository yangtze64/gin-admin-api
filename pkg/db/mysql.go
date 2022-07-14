package db

import (
	"fmt"
	"gin-admin-api/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strings"
	"sync"
)

type MysqlPool struct {
	MaxConn int
	MaxIdle int
}

type MysqlRw struct {
	DataSource string
}

type MySqlConf struct {
	ConnName   string
	Host       string
	Port       int
	User       string
	Pwd        string
	Database   string
	Charset    string
	DataSource string
	prefix     string
	Read       MysqlRw
	Write      MysqlRw
	Pool       MysqlPool
}

type Conn struct {
	dbr  *gorm.DB
	dbw  *gorm.DB
	name string
}

type Databases struct {
	Conns []Conn
	Names map[string]int
}

var (
	once sync.Once
	Dbs  *Databases
)

func NewMysql(dbConf []MySqlConf) *Databases {
	once.Do(func() {
		var (
			connArr = make([]Conn, 0)
			nameArr = make(map[string]int)
		)
		if len(dbConf) > 0 {
			for _, c := range dbConf {
				if c.ConnName == "" {
					logger.Fatalf("Db conn name can't be empty")
				}
				l := len(nameArr)
				nameArr[c.ConnName] = 1
				if len(nameArr) == l {
					logger.Fatalf("Db conn name can not repeat")
				}
			}
		} else {
			logger.Fatalf("Db none config info")
		}

		var (
			dbr      *gorm.DB
			dbw      *gorm.DB
			dbConfig *gorm.Config
			dsn      string
		)
		for i, c := range dbConf {
			dbConfig = &gorm.Config{
				NamingStrategy: schema.NamingStrategy{
					TablePrefix:   c.prefix,
					SingularTable: true,
				},
			}
			dsn = c.DataSource
			if dsn != "" {
				dbr = setOrmInstancePool(getOrmInstance(dsn, dbConfig), c.Pool)
				dbw = dbr
			} else {
				if c.Read.DataSource != "" {
					dbr = setOrmInstancePool(getOrmInstance(c.Read.DataSource, dbConfig), c.Pool)
				}
				if c.Write.DataSource != "" {
					dbw = setOrmInstancePool(getOrmInstance(c.Write.DataSource, dbConfig), c.Pool)
				}
				if dbr == nil && dbw == nil {
					dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=PRC", c.User, c.Pwd, c.Host, c.Port, c.Database, c.Charset)
					dbr = setOrmInstancePool(getOrmInstance(dsn, dbConfig), c.Pool)
					dbw = dbr
				} else {
					if dbr == nil {
						dbr = dbw
					}
					if dbw == nil {
						dbw = dbr
					}
				}
			}
			conn := Conn{dbr: dbr, dbw: dbw, name: c.ConnName}
			if c.ConnName == "default" && i > 0 {
				connArr = append([]Conn{conn}, connArr...)
			} else {
				connArr = append(connArr, conn)
			}
		}
		Dbs = &Databases{Conns: connArr, Names: nameArr}
	})
	return Dbs
}

func Orm() *gorm.DB {
	return OrmR()
}

func OrmR() *gorm.DB {
	db := GetDbs()
	return db.Conns[0].dbr
}

func OrmW() *gorm.DB {
	db := GetDbs()
	return db.Conns[0].dbw
}

func GetOrm(v ...string) *gorm.DB {
	var orm *gorm.DB
	dbs := GetDbs()
	len := len(v)
	if len > 0 {
		_, ok := dbs.Names[v[0]]
		if !ok {
			logger.Panicf("NO Exist %s Connect", v[0])
		}
		for _, conn := range dbs.Conns {
			if conn.name == v[0] {
				if len > 1 {
					rw := strings.ToLower(v[1])
					if rw == "r" {
						orm = conn.dbr
					} else if rw == "w" {
						orm = conn.dbw
					} else {
						logger.Panicf("NO Exist %s.%s Connect,Use `r` or `w` to get the instance", v[0], v[1])
					}
				} else {
					orm = conn.dbr
				}
			}
		}
	} else {
		orm = dbs.Conns[0].dbr
	}
	return orm
}

func GetDbs() *Databases {
	if Dbs == nil {
		logger.Panicf("Dbs Unset instance")
	}
	return Dbs
}

// 获取gorm实例
func getOrmInstance(dsn string, gc *gorm.Config) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), gc)
	if err != nil {
		logger.Fatal(err)
	}
	return db
}

func setOrmInstancePool(db *gorm.DB, pool MysqlPool) *gorm.DB {
	dbConn, _ := db.DB()
	if pool.MaxConn > 0 {
		dbConn.SetMaxOpenConns(pool.MaxConn)
	}
	if pool.MaxIdle > 0 {
		dbConn.SetMaxIdleConns(pool.MaxIdle)
	}
	return db
}
