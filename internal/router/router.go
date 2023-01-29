package router

import (
	"gin-admin-api/internal/handler"
	"gin-admin-api/internal/svc"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, serverCtx *svc.ServiceContext) {
	r.GET("/", handler.IndexHandler(serverCtx))
	api := r.Group("/api")
	PassportRouter(api, serverCtx)
}
