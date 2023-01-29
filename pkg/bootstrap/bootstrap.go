package bootstrap

import (
	"flag"
	"gin-admin-api/internal/config"
	"gin-admin-api/internal/router"
	"gin-admin-api/internal/svc"
	"gin-admin-api/pkg/conf"
	"gin-admin-api/pkg/logx"
	"gin-admin-api/pkg/server"
)

var (
	configFile = flag.String("f", "etc/config.yaml", "The Config File")
	c          config.Config
	srv        *server.Server
)

func init() {
	flag.Parse()
	conf.MustLoad(*configFile, &c)
	// fmt.Printf("%#+v\n", c)
	logx.Setup(c.Log)
	// logx.WithField("ServerName", c.Server.Name)

	srv = server.MustNewServer(c.Server)

	ctx := svc.NewServiceContext(c)

	router.Register(srv.GetEngine(), ctx)
}

func Run() {
	srv.Run()
}
