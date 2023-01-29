package db

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

// type DB struct {
// 	conf  map[string]*MySqlConf
// 	Conns map[string]*gorm.DB
// }

var (
	dbOnce      sync.Once
	configs     map[string]*MySqlConf
	Conns       map[string]*gorm.DB
	defaultName = "default"
)

func Setup(config map[string]*MySqlConf) {
	lens := len(config)
	if lens <= 0 {
		panic("DB configuration not set")
	}
	dbOnce.Do(func() {
		conns := make(map[string]*gorm.DB, lens)
		for name, conf := range config {
			var gdb *gorm.DB
			if conf.Write != "" || len(conf.Read) > 0 {
				gdb = connectRWDB(conf)
			} else {
				gdb = connectDB(conf)
			}
			conns[strings.TrimSpace(name)] = gdb
		}
		Conns = conns
		configs = config
	})
}

func connectDB(c *MySqlConf) *gorm.DB {
	dsn := FormatConn(c)
	db := getInstance(dsn, func() *gorm.Config {
		return getOrmConf(c)
	})
	return setOrmPool(db, c.Pool)
}

func connectRWDB(c *MySqlConf) *gorm.DB {
	write := c.Write
	read := c.Read
	dsn := FormatConn(c)
	if write == "" {
		write = dsn
	}
	if len(read) == 0 {
		read = append(read, dsn)
	}
	db := getInstance(write, func() *gorm.Config {
		return getOrmConf(c)
	})

	replicas := []gorm.Dialector{}
	for _, s := range read {
		dialector := getOrmDialector(s)
		replicas = append(replicas, dialector)
	}
	resolver := dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{getOrmDialector(write)},
		Replicas: replicas,
		Policy:   dbresolver.RandomPolicy{},
	})
	if c.Pool.MaxConn > 0 {
		resolver.SetMaxOpenConns(c.Pool.MaxConn)
	}
	if c.Pool.MaxIdle > 0 {
		resolver.SetMaxIdleConns(c.Pool.MaxIdle)
	}
	db.Use(resolver)
	return db
}

func FormatConn(c *MySqlConf) string {
	if c.DataSource != "" {
		return c.DataSource
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=PRC", c.User, c.Pwd, c.Host, c.Port, c.Database, c.Charset)
	return dsn
}

func getOrmConf(c *MySqlConf) *gorm.Config {
	return &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.Prefix,
			SingularTable: true,
		},
	}
}

func getOrmDialector(dsn string) gorm.Dialector {
	return mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	})
}

// 获取gorm实例
func getInstance(dsn string, fns ...func() *gorm.Config) *gorm.DB {
	var gc *gorm.Config
	for _, fn := range fns {
		gc = fn()
	}
	db, err := gorm.Open(getOrmDialector(dsn), gc)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func setOrmPool(db *gorm.DB, pool MysqlPool) *gorm.DB {
	dbConn, _ := db.DB()
	if pool.MaxConn > 0 {
		dbConn.SetMaxOpenConns(pool.MaxConn)
	}
	if pool.MaxIdle > 0 {
		dbConn.SetMaxIdleConns(pool.MaxIdle)
	}
	return db
}

func GetConn(name ...string) *gorm.DB {
	key := defaultName
	if len(name) > 0 {
		key = name[0]
	}
	if db, ok := Conns[key]; ok {
		return db
	}
	panic("connect with name not exist")
}
