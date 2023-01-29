package svc

import (
	"gin-admin-api/internal/config"
	"gin-admin-api/internal/dao"
	"gin-admin-api/pkg/db"
)

type ServiceContext struct {
	Config config.Config
	// Db     *gorm.DB
	// UserDb *gorm.DB
	UserDao dao.IUserDao
}

func NewServiceContext(c config.Config) *ServiceContext {
	db.Setup(c.DBs)
	conn := db.GetConn()
	return &ServiceContext{
		Config: c,
		// Db:     db.GetConn(),
		// Db:     db.GetConn("default"),
		// UserDb: db.GetConn("user"),
		UserDao: dao.User(conn),
	}
}
