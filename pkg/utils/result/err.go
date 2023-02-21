package result

import "fmt"

type Err struct {
	Bean
}

func (e *Err) Error() string {
	return fmt.Sprintf("ErrCode:%d,ErrMsg:%s", e.GetCode(), e.GetMsg())
}

func NewErr(code CodeType, msg string, data ...interface{}) *Err {
	e := &Err{}
	e.Code = code
	e.Msg = msg
	if len(data) > 0 {
		e.Data = data[0]
	}
	return e
}
