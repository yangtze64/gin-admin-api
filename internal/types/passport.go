package types

type (
	// SignupReq 注册 请求参数
	SignupReq struct {
		MobileReq
		Username       string `json:"username" form:"username" binding:"required,gte=4,lte=30"`
		Password       string `json:"password" form:"password" binding:"required,gte=6,lte=20"`
		PasswordRepeat string `json:"passwordRepeat" form:"passwordRepeat" binding:"required,eqfield=Password"`
		Email          string `json:"email" form:"email" binding:"required,email"`
		Weixin         string `json:"weixin" form:"weixin" binding:"omitempty,gte=6,lte=20"`
	}
)
