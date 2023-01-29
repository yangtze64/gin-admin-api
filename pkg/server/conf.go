package server

type (
	Conf struct {
		Name     string `default:"httpServer" required:"true"`
		Mode     string `default:"dev"`
		Host     string `default:"0.0.0.0"`
		Port     int    `default:"8080"`
		CertFile string
		KeyFile  string
	}
)