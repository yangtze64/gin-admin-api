package shared

import (
	"gin-admin-api/pkg/errx"
	"net/http"
)

const (
	SignatureError errx.CodeType = 200001
)

var (
	codemsg = map[errx.CodeType]string{
		errx.OK:                        "success",
		errx.StatusInternalServerError: "Server Error",

		errx.DBError:       "Database Busy",
		errx.BusinessError: "Operation Exception",
		SignatureError:     "Signature Error",
	}
)

func GetMessage(code errx.CodeType) string {
	if msg, ok := codemsg[code]; ok {
		return msg
	}
	if msg := http.StatusText(int(code)); msg != "" {
		return msg
	}
	return ""
}
