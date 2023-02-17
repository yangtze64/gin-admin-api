package server

type (
	Conf struct {
		Name     string `default:"httpServer" required:"true"`
		Mode     string `default:"dev" options:"dev,test,prod"`
		Host     string `default:"0.0.0.0"`
		Port     int    `default:"8080"`
		Pprof    bool   `default:"false"`
		CertFile string
		KeyFile  string
	}
)
