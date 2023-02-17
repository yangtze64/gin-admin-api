package bootstrap

import (
	"flag"
	"gin-admin-api/internal/router"
	"gin-admin-api/internal/svc"
	"gin-admin-api/internal/utils/global"
	"gin-admin-api/pkg/conf"
	"gin-admin-api/pkg/logx"
	"gin-admin-api/pkg/server"
	"github.com/sirupsen/logrus"
)

var (
	configFile = flag.String("f", "etc/config.yaml", "The Config File")
	srv        *server.Server
)

func init() {
	flag.Parse()
	c := &global.C
	conf.MustLoad(*configFile, c)
	// fmt.Printf("%#+v\n", c)
	logx.Setup(c.Log)
	logx.AddGlobalField("ServerName", c.Server.Name)
	logx.AddGlobalFields(logrus.Fields{"ServerName1": c.Server.Name})

	srv = server.MustNewServer(c.Server)

	ctx := svc.NewServiceContext(*c)

	router.Register(srv.GetEngine(), ctx)
}

func Run() {
	srv.Run()
}
