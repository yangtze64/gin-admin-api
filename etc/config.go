package etc

import (
	"gin-admin-api/pkg/cache"
	"gin-admin-api/pkg/db"
	"gin-admin-api/pkg/logger"
	"gin-admin-api/pkg/server"
)

type Config struct {
	Server server.ServeConf
	Dbs    []db.MySqlConf
	Redis  cache.RedisConf
	Log    logger.LogConf
	Cors   struct {
		Domain []string
	}
}
