package pkg

import (
	"flag"
	"fmt"
	"gin-admin-api/pkg/cache"
	"gin-admin-api/pkg/config"
	"gin-admin-api/pkg/db"
	"gin-admin-api/pkg/logger"
	"gin-admin-api/pkg/server"
	"gin-admin-api/pkg/utils/console"
)

var configFile = flag.String("f", "etc/config.yaml", "The Config File")
var packStatic = flag.Bool("pks", false, "Pack Static File")

func Bootstrap() {
	flag.Parse()
	fmt.Println(console.Green(fmt.Sprintf("Load Config File: %s\n", *configFile)))
	config.LoadConfig(*configFile, *packStatic)
	Initialize()
}

func Initialize() {
	logger.Setup(config.Config.Log)
	server.New(config.Config.Server)
	db.NewMysql(config.Config.Dbs)
	cache.NewRedis(config.Config.Redis)
}
