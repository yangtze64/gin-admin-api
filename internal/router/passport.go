package router

import (
	"gin-admin-api/internal/handler/passport"
	"gin-admin-api/internal/svc"

	"github.com/gin-gonic/gin"
)

func PassportRouter(r *gin.RouterGroup, serverCtx *svc.ServiceContext) {
	v1 := r.Group("/v1")
	{
		v1.POST("/signup", passport.SignupHandler(serverCtx))
	}
}
