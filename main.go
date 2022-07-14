package main

import (
	"gin-admin-api/app/router"
	"gin-admin-api/pkg"
	"gin-admin-api/pkg/server"
)

func init() {
	server.Routers = append(server.Routers, router.InitRouter)
}

func main() {
	pkg.Bootstrap()
	server.Srv.Run()
}
