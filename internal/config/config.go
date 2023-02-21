package config

import (
	"gin-admin-api/pkg/db"
	"gin-admin-api/pkg/logx"
	"gin-admin-api/pkg/redisclient"
	"gin-admin-api/pkg/server"
)

type Config struct {
	Server server.Conf
	Log    logx.Conf
	DBs    map[string]*db.MySqlConf
	Redis  redisclient.Conf
}
