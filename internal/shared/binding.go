package shared

import (
	"gin-admin-api/pkg/utils"
	"github.com/gin-gonic/gin"
)

func ShouldBind(ctx *gin.Context, v interface{}) error {
	if err := utils.SetStructValue(v); err != nil {
		return err
	}
	if err := ctx.ShouldBind(v); err != nil {
		return err
	}
	return nil
}
