package shared

import (
	"gin-admin-api/internal/utils/global"
	"gin-admin-api/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShouldBind(ctx *gin.Context, v interface{}) error {
	if err := utils.SetStructValue(v); err != nil {
		if global.IsProd() {
			ctx.PureJSON(http.StatusBadRequest, gin.H{"msg": "参数错误"})
		} else {
			ctx.PureJSON(http.StatusBadRequest, gin.H{"msg": err})
		}
		return err
	}
	if err := ctx.ShouldBind(v); err != nil {
		if global.IsProd() {
			ctx.PureJSON(http.StatusBadRequest, gin.H{"msg": "参数校验错误"})
		} else {
			ctx.PureJSON(http.StatusBadRequest, gin.H{"msg": err})
		}
		return err
	}
	return nil
}
