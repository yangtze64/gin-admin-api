package logic

import "github.com/gin-gonic/gin"

type IMember interface {
	GetMember(ctx *gin.Context)
}

var memberImpl IMember

func Member() IMember {
	return memberImpl
}

func RegisterMember(i IMember) {
	memberImpl = i
}
