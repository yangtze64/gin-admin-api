package logic

import "gin-admin-api/internal/types"

var passportLogic IPassportLogic

type IPassportLogic interface {
	// Signup 注册用户
	Signup(req *types.SignupReq) (resp *types.EmptyResp, err error)
}

func RegisterPassportLogic(i IPassportLogic) {
	passportLogic = i
}
