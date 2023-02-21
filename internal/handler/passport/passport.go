package passport

import (
	"gin-admin-api/internal/logic"
	"gin-admin-api/internal/svc"
	"gin-admin-api/internal/types"
	"gin-admin-api/internal/utils/httpx"
	"github.com/gin-gonic/gin"
)

func SignupHandler(svcCtx *svc.ServiceContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req types.SignupReq
		if !httpx.CheckParamsFailRenderJson(ctx, &req) {
			return
		}
		l := logic.Passport(ctx, svcCtx)
		resp, err := l.Signup(&req)
		httpx.Json(ctx, resp, err)
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
