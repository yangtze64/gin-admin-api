package logic

import (
	"context"
	"gin-admin-api/internal/logic/impl/passport"
	"gin-admin-api/internal/svc"
)

func Passport(ctx context.Context, svcCtx *svc.ServiceContext) IPassportLogic {
	if passportLogic == nil {
		RegisterPassportLogic(passport.NewPassportLogic(ctx, svcCtx))
	}
	return passportLogic
}
