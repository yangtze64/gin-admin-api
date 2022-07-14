package server

import (
	"context"
	"fmt"
	"gin-admin-api/pkg/logger"
	"gin-admin-api/pkg/utils"
	"gin-admin-api/pkg/utils/console"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

// 环境模式
const (
	DevMode  = "dev"
	TestMode = "test"
	ProMode  = "prod"
)

type ServeConf struct {
	Name      string
	Host      string
	Port      int
	Env       string
	EnableSsl bool
	CertFile  string
	KeyFile   string
}

type Server struct {
	Cfg    ServeConf
	Engine *gin.Engine
	mutex  sync.Mutex
}

var (
	Srv     *Server
	Routers = make([]func(r *gin.Engine), 0)
)

func New(cfg ServeConf) *Server {
	Srv = &Server{Cfg: cfg}
	Srv.WithEngine().
		InitRouter()
	return Srv
}

func (s *Server) InitRouter() {
	s.GlobalMiddleware()
	for _, f := range Routers {
		f(s.Engine)
	}
}

func (s *Server) WithEngine() *Server {
	if s.Engine == nil {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		if s.Cfg.Env == ProMode {
			gin.SetMode(gin.ReleaseMode)
		}
		engine := gin.New()
		s.Engine = engine
	}
	return s
}

func (s *Server) GetEngine() *gin.Engine {
	return s.Engine
}

func (s *Server) Run() {
	addr := fmt.Sprintf("%s:%d", s.Cfg.Host, s.Cfg.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: s.Engine,
	}
	go func() {
		// 服务连接
		if s.Cfg.EnableSsl {
			if err := srv.ListenAndServeTLS(s.Cfg.CertFile, s.Cfg.KeyFile); err != nil && err != http.ErrServerClosed {
				logger.Fatalf("listen: %s\n", err)
			}
		} else {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logger.Fatalf("listen: %s\n", err)
			}
		}
	}()
	logger.Infof("Listen %s", addr)
	fmt.Println(console.Green("Server run at:"))
	fmt.Printf("-  Local:   http://localhost:%d/ \r\n", s.Cfg.Port)
	fmt.Printf("-  Network: http://%s:%d/ \r\n", utils.GetLocalHost(), s.Cfg.Port)
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Server Shutdown Error:", err)
	}
	logger.Info("Server Shutdown")
}
