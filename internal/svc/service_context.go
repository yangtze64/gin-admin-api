package svc

import (
	"gin-admin-api/internal/config"
	"gin-admin-api/internal/dao"
	"gin-admin-api/pkg/db"
)

type ServiceContext struct {
	Config  config.Config
	UserDao dao.IUserDao
}

func NewServiceContext(c config.Config) *ServiceContext {
	db.Setup(c.DBs)
	conn := db.GetConn()
	return &ServiceContext{
		Config:  c,
		UserDao: dao.User(conn),
	}
}
