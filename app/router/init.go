package router

import (
	"gin-admin-api/app/api"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.GET("/member", api.Member.GetMemberList)
}
