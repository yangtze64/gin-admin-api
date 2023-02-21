package httpx

import (
	"gin-admin-api/internal/utils/shared"
	"gin-admin-api/pkg/errx"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func BeanJson(ctx *gin.Context, bean errx.Beaner) {
	code := bean.GetCode()
	if bean.GetCode() == 0 {
		code = errx.StatusBadRequest
		bean.SetCode(code)
	} else if http.StatusText(code) == "" {
		code = errx.StatusBadRequest
	}
	if bean.GetMsg() == "" {
		bean.SetMsg(shared.GetMessage(bean.GetCode()))
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

func Json(ctx *gin.Context, data interface{}, err error) {
	if err == nil {
		SuccessJson(ctx, data)
	} else {
		causeErr := errors.Cause(err) // err类型
		if e, ok := causeErr.(errx.Beaner); ok {
			BeanJson(ctx, e)
		} else {
			ErrorJson(ctx, err.Error(), errx.StatusBadRequest)
		}
	}
}
