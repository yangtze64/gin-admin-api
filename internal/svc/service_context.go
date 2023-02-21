package svc

import (
	"gin-admin-api/internal/config"
	"gin-admin-api/internal/dao"
	"gin-admin-api/pkg/db"
	"gin-admin-api/pkg/redisclient"
)

type ServiceContext struct {
	Config  config.Config
	Redis   *redisclient.RedisClient
	UserDao dao.IUserDao
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisClient := redisclient.New(c.Redis)
	db.Setup(c.DBs)
	conn := db.GetConn()
	return &ServiceContext{
		Config:  c,
		Redis:   redisClient,
		UserDao: dao.User(conn),
	}
}
