package conf

import (
	"gin-admin-api/pkg/utils"
)

// 解析预制配置
func presetConf(v interface{}) error {
	return utils.SetStructValue(v,
		utils.WithSetStructFieldDefault(),
		utils.WithVerifyStructFieldRequired(),
		utils.WithVerifyStructFieldRange(),
		utils.WithVerifyStructFieldOptions(),
	)
}
