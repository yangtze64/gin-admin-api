package result

type (
	Beaner interface {
		GetCode() CodeType
		SetCode(code CodeType)
		GetMsg() string
		SetMsg(msg string)
		GetData() interface{}
		SetData(data interface{})
		i()
	}

	Bean struct {
		Code CodeType    `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}
)

func (b Bean) i() {}

func (b Bean) GetCode() CodeType {
	return b.Code
}

func (b Bean) SetCode(code CodeType) {
	b.Code = code
}

func (b Bean) GetMsg() string {
	return b.Msg
}

func (b Bean) SetMsg(msg string) {
	b.Msg = msg
}

func (b Bean) GetData() interface{} {
	return b.Data
}

func (b Bean) SetData(data interface{}) {
	b.Data = data
}
