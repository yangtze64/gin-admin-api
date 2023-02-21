package shared

import (
	"gin-admin-api/pkg/result"
	"net/http"
)

const (
	SignatureError result.CodeType = 200001
)

var (
	codemsg = map[result.CodeType]string{
		result.OK:                        "success",
		result.StatusInternalServerError: "Server Error",

		result.DBError:       "Database Busy",
		result.BusinessError: "Operation Exception",
		SignatureError:       "Signature Error",
	}
)

func GetMessage(code result.CodeType) string {
	if msg, ok := codemsg[code]; ok {
		return msg
	}
	if msg := http.StatusText(int(code)); msg != "" {
		return msg
	}
	return ""
}
