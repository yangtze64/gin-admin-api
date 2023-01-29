package logic

var passportLogic IPassportLogic

type IPassportLogic interface {
	GetMember()
}

func RegisterPassportLogic(i IPassportLogic) {
	passportLogic = i
}
