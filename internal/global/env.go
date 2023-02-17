package global

import "gin-admin-api/pkg/utils"

func IsDev() bool {
	return C.Server.Mode == utils.DevMode
}

func IsTest() bool {
	return C.Server.Mode == utils.TestMode
}

func IsProd() bool {
	return C.Server.Mode == utils.ProdMode
}
