package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

type (
	Server struct {
		srv    *http.Server
		engine *gin.Engine
		conf   Conf
	}
)

func MustNewServer(c Conf, opts ...RunOption) *Server {
	server, err := NewServer(c, opts...)
	if err != nil {
		log.Fatal(err)
	}
	return server
}

func NewServer(c Conf, opts ...RunOption) (*Server, error) {
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	engine := gin.New()
	fmt.Println(c)
	if c.Pprof {
		pprof.Register(engine)
	}
	srv := &http.Server{
		Addr:    addr,
		Handler: engine,
	}
	server := &Server{
		srv:    srv,
		engine: engine,
		conf:   c,
	}
	// opts = append([]RunOption{WithNotFoundHandler(nil)}, opts...)
	for _, opt := range opts {
		opt(server)
	}

	return server, nil
}

func (s *Server) Run() {
	go s.Start()
	s.StopSignal()
}

func (s *Server) Start() {
	// 服务连接
	if s.conf.CertFile != "" && s.conf.KeyFile != "" {
		if err := s.srv.ListenAndServeTLS(s.conf.CertFile, s.conf.KeyFile); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	} else {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}
}

func (s *Server) StopSignal() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.Stop(ctx)
}

func (s *Server) Stop(ctx context.Context) {
	if err := s.srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}

func (s *Server) GetEngine() *gin.Engine {
	return s.engine
}

func (s *Server) GetSrv() *http.Server {
	return s.srv
}
