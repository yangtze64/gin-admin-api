package handler

import (
	"gin-admin-api/internal/svc"
	"github.com/gin-gonic/gin"
)

func IndexHandler(svcCtx *svc.ServiceContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.String(200, "Hi gin-admin-api")
	}
}
