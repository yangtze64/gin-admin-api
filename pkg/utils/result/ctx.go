package result

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BeanJson(ctx *gin.Context, bean Beaner) {
	code := bean.GetCode()
	if http.StatusText(code) == "" {
		code = http.StatusOK
	}
	ctx.PureJSON(code, bean)
}

func Success(ctx *gin.Context, data interface{}) {
	bean := NewErr(OK, "", data)
	BeanJson(ctx, bean)
}

func Error(ctx *gin.Context, msg string, code CodeType, data ...interface{}) {
	bean := NewErr(code, msg, data...)
	BeanJson(ctx, bean)
}
