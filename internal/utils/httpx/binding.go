package httpx

import (
	"gin-admin-api/internal/utils/global"
	"gin-admin-api/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ShouldBind 绑定参数
func ShouldBind(ctx *gin.Context, v interface{}) error {
	// 预设default默认值
	if err := utils.SetStructValue(v); err != nil {
		return err
	}
	if err := ctx.ShouldBind(v); err != nil {
		return err
	}
	return nil
}

// CheckParamsFailRenderJson 校验参数失败write json 校验通过返回true
func CheckParamsFailRenderJson(ctx *gin.Context, v interface{}) bool {
	if err := ShouldBind(ctx, v); err != nil {
		if global.IsProd() {
			ErrorJson(ctx, "Parameter Error", http.StatusBadRequest)
		} else {
			ErrorJson(ctx, err.Error(), http.StatusBadRequest)
		}
		return false
	}
	return true
}
