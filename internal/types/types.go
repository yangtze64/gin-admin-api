package types

type (
	EmptyReq  struct{}
	EmptyResp struct{}
	PagerReq  struct {
		Page  int `json:"page" form:"page" default:"1" binding:"required,min=1"`
		Limit int `json:"limit" form:"limit" default:"10" binding:"required,min=1,max=100"`
	}
	MobileReq struct {
		Mobile   string `json:"mobile" form:"mobile" binding:"required,gte=4,lte=20"`
		Areacode int    `json:"areacode" form:"areacode" default:"86" binding:"required,max=9999"`
	}
)
