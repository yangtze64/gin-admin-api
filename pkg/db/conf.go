package db

type (
	MySqlConf struct {
		ConnName   string
		Host       string
		Port       int `default:"3306"`
		User       string
		Pwd        string
		Database   string
		Charset    string `default:"utf8mb4"`
		DataSource string
		Prefix     string
		Write      string
		Read       []string
		Pool       MysqlPool
	}

	MysqlPool struct {
		MaxConn int
		MaxIdle int
	}
)
