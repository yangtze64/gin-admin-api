package server

import (
	"gin-admin-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (s *Server) GlobalMiddleware() {
	s.healthRouter()
	s.Engine.Use(CORS, RequestId, AccessLog)
}

func (s *Server) healthRouter() {
	s.Engine.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "")
	})
}

func CORS(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}
	ctx.Next()
}

func RequestId(ctx *gin.Context) {
	reqId := ctx.Request.Header.Get("X-Request-Id")
	if reqId == "" {
		reqId = uuid.New().String()
		ctx.Request.Header.Set("X-Request-Id", reqId)
	}
	// 写入响应
	ctx.Header("X-Request-Id", reqId)
	ctx.Next()
}

func AccessLog(ctx *gin.Context) {
	reqId := ctx.Request.Header.Get("X-Request-Id")
	log := logger.GetLogger().WithField("X-Request-Id", reqId).WithField("IP", ctx.ClientIP())
	ctx.Set("log", log)
	ctx.Next()
}
