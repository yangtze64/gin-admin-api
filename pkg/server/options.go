package server

type (
	RunOption func(*Server)
)

func WithRegisterOnShutdown(fun func()) RunOption {
	return func(s *Server) {
		s.srv.RegisterOnShutdown(fun)
	}
}
