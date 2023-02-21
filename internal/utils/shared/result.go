package shared

import (
	"gin-admin-api/pkg/errx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BeanJson(ctx *gin.Context, bean errx.Beaner) {
	code := bean.GetCode()
	if http.StatusText(code) == "" {
		code = http.StatusOK
	}
	ctx.PureJSON(code, bean)
}

func SuccessJson(ctx *gin.Context, data interface{}) {
	BeanJson(ctx, Success(data))
}

func ErrorJson(ctx *gin.Context, msg string, code errx.CodeType, data ...interface{}) {
	BeanJson(ctx, Error(msg, code, data...))
}

func Success(data interface{}) *errx.Err {
	bean := errx.NewErr(errx.OK, "", data)
	return bean
}

func Error(msg string, code errx.CodeType, data ...interface{}) *errx.Err {
	bean := errx.NewErr(code, msg, data...)
	return bean
}
