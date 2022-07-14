package member

import "github.com/gin-gonic/gin"

type memberLogic struct {
}

func New() *memberLogic {
	m := &memberLogic{}
	return m
}

func (m *memberLogic) GetMember(ctx *gin.Context) {

}
