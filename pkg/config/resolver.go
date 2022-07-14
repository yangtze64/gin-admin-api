package config

import (
	"bytes"
	"gin-admin-api/etc"
	"gin-admin-api/pkg/utils"
	"github.com/spf13/viper"
	"log"
	"path"
	"strings"
	"sync"
)

var (
	once     sync.Once
	resolver *Resolver
	Config   etc.Config
)

type Resolver struct {
	V    *viper.Viper
	File string
}

func NewResolver(confPath string, isPackStatic bool) (*Resolver, error) {
	once.Do(func() {
		v := viper.New()
		resolver = &Resolver{V: v, File: confPath}
		if isPackStatic {
			ext := path.Ext(confPath)
			configBytes, err := etc.Asset(confPath)
			if err != nil {
				log.Fatalf("Asset() can not found setting file,%s", err.Error())
			}
			//设置要读取的文件类型
			resolver.V.SetConfigType(strings.Trim(ext, "."))
			//读取
			err = resolver.V.ReadConfig(bytes.NewBuffer(configBytes))
			if err != nil {
				log.Fatalf("ReadConfig() can not read file,%s", err.Error())
			}
		} else {
			resolver.V.SetConfigFile(confPath)
			if err := resolver.V.ReadInConfig(); err != nil {
				log.Fatalf(err.Error())
			}
		}
	})
	return resolver, nil
}

func LoadConfig(file string, isPackStatic bool, v ...interface{}) {
	if !isPackStatic {
		if ok := utils.FileIsExist(file); !ok {
			log.Fatalf("error: config file %s, %s", file, "file does not exist")
		}
	}
	newResolver, err := NewResolver(file, isPackStatic)
	if err != nil {
		log.Fatalf("error: config file %s, %s", file, err.Error())
	}
	if msg := newResolver.V.Unmarshal(&Config); msg != nil {
		log.Fatalf("error: config file %s, %s", file, msg.Error())
	}
	if len(v) > 0 {
		v[0] = Config
	}
}
