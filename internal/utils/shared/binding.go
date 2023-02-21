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
			ErrorJson(ctx, "Parameter Error", http.StatusBadRequest)
		} else {
			ErrorJson(ctx, err.Error(), http.StatusBadRequest)
		}
		return err
	}
	if err := ctx.ShouldBind(v); err != nil {
		if global.IsProd() {
			ErrorJson(ctx, "Parameter Check Error", http.StatusBadRequest)
		} else {
			ErrorJson(ctx, err.Error(), http.StatusBadRequest)
		}
		return err
	}
	return nil
}
