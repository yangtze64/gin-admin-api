package redisclient

type Conf struct {
	Host     string
	Port     int `default:"6379"`
	Auth     string
	Db       int `default:"0"`
	MaxConn  int
	MaxIdle  int
	RetryNum int
}
