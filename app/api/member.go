package api

import (
	"gin-admin-api/app/logic/dto"
	"gin-admin-api/pkg/base"
	"gin-admin-api/pkg/logger"
	"gin-admin-api/pkg/utils"
	"github.com/gin-gonic/gin"
)

type memberApi struct {
	base.Api
}

var Member = &memberApi{}

func (m *memberApi) GetMemberList(ctx *gin.Context) {
	err := m.Bind(ctx, &dto.Pager{}).Err
	if err != nil {
		logger.Errorf(err.Error())
	}
	logz := utils.GetLog(ctx)
	logz.Infof("abcd %s", "哈哈")
	// logic.Member().GetMember(ctx)
}
