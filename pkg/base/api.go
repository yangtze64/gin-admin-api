package base

import (
	"fmt"
	"gin-admin-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Api struct {
	Err error
}

func (a *Api) AddErr(err error) {
	if a.Err == nil {
		a.Err = err
	} else if err != nil {
		a.Err = fmt.Errorf("%v; %w", a.Err, err)
	}
}

// Bind 参数校验
func (a *Api) Bind(c *gin.Context, d interface{}, bindings ...binding.Binding) *Api {
	var err error
	if len(bindings) == 0 {
		bindings = constructor.GetBindingForGin(d)
	}
	for i := range bindings {
		if bindings[i] == nil {
			err = c.ShouldBindUri(d)
		} else {
			err = c.ShouldBindWith(d, bindings[i])
		}
		if err != nil && err.Error() == "EOF" {
			logger.Warning("request body is not present anymore. ")
			err = nil
			continue
		}
		if err != nil {
			a.AddErr(err)
			break
		}
	}
	return a
}
