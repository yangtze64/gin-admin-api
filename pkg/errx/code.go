package errx

import "net/http"

type CodeType = int

const (
	// 常用httpCode
	OK                        CodeType = http.StatusOK
	StatusBadRequest          CodeType = http.StatusBadRequest
	StatusUnauthorized        CodeType = http.StatusUnauthorized
	StatusForbidden           CodeType = http.StatusForbidden
	StatusNotFound            CodeType = http.StatusNotFound
	StatusRequestTimeout      CodeType = http.StatusRequestTimeout
	StatusInternalServerError CodeType = http.StatusInternalServerError

	// DBError 数据库错误
	DBError CodeType = 100000
	// BusinessError 业务错误领域
	BusinessError CodeType = 200000
)
