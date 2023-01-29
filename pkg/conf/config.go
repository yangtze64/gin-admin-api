package conf

import (
	"log"

	"github.com/spf13/viper"
)

var Resolver *viper.Viper

func Load(file string, v interface{}) error {
	resolver := viper.New()
	resolver.SetConfigFile(file)
	if err := resolver.ReadInConfig(); err != nil {
		return err
	}
	if err := resolver.Unmarshal(v); err != nil {
		return err
	}
	Resolver = resolver
	err := presetConf(v)
	return err
}

func MustLoad(path string, v interface{}) {
	if err := Load(path, v); err != nil {
		log.Fatalf("error: config file %s, %s", path, err.Error())
	}
}
