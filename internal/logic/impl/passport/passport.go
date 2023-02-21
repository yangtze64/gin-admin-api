package passport

import (
	"context"
	"gin-admin-api/internal/svc"
	"gin-admin-api/internal/types"
	"gin-admin-api/pkg/logx"
)

type PassportLogic struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPassportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PassportLogic {
	return &PassportLogic{
		logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (p *PassportLogic) Signup(req *types.SignupReq) (resp *types.EmptyResp, err error) {
	// fmt.Println(req)
	// err = httpx.Error("用户不存在", errx.StatusNotFound)
	return
}
