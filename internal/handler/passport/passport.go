package passport

import (
	"fmt"
	"gin-admin-api/internal/logic"
	"gin-admin-api/internal/svc"
	"gin-admin-api/internal/types"
	"gin-admin-api/internal/utils/shared"
	"github.com/gin-gonic/gin"
)

func SignupHandler(svcCtx *svc.ServiceContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req types.SignupReq
		if err := shared.ShouldBind(ctx, &req); err != nil {
			return
		}
		l := logic.Passport(ctx, svcCtx)
		resp, err := l.Signup(&req)
		fmt.Printf("%#+v\n", err)
		fmt.Printf("%#+v\n", resp)
	}
}

func LoginHandler(svcCtx *svc.ServiceContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

	}
}

func LogoutHandler(svcCtx *svc.ServiceContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

	}
}
