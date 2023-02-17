package global

import (
	"encoding/json"
	"gin-admin-api/internal/config"
)

var C config.Config

func ConfigMap() (map[string]interface{}, error) {
	var m map[string]interface{}
	b, err := json.Marshal(&C)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func Config(key string) interface{} {
	return nil
}
