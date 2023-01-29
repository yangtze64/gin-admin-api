package passport

import (
	"gin-admin-api/internal/logic"
	"gin-admin-api/internal/svc"

	"github.com/gin-gonic/gin"
)

func LoginHandler(svcCtx *svc.ServiceContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		l := logic.Passport(ctx, svcCtx)
		l.GetMember()
	}
}
