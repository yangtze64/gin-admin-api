package cache

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
)

type RedisConf struct {
	Host     string
	Port     int
	Auth     string
	Db       int
	MaxConn  int
	MaxIdle  int
	RetryNum int
}

var (
	redisOnce sync.Once
	Redis     *redis.Client
)

func NewRedis(c RedisConf) *redis.Client {
	redisOnce.Do(func() {
		options := &redis.Options{
			Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
			Password: c.Auth,
			DB:       c.Db,
		}
		if c.MaxConn > 0 {
			options.PoolSize = c.MaxConn
		}
		if c.MaxIdle > 0 {
			options.MinIdleConns = c.MaxIdle
		}
		if c.RetryNum > 0 {
			options.MaxRetries = c.RetryNum
		}
		Redis = redis.NewClient(options)
	})
	return Redis
}
