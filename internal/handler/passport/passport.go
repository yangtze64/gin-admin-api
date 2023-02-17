package passport

import (
	"fmt"
	"gin-admin-api/internal/logic"
	"gin-admin-api/internal/shared"
	"gin-admin-api/internal/svc"
	"gin-admin-api/internal/types"

	"github.com/gin-gonic/gin"
)

func SignupHandler(svcCtx *svc.ServiceContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var (
			req  types.SignupReq
			resp *types.EmptyResp
		)
		if err := shared.ShouldBind(ctx, &req); err != nil {
			l := logic.Passport(ctx, svcCtx)
			resp, err = l.Signup(&req)
			fmt.Printf("%#+v\n", resp)
		}
	}
}
