package redisclient

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	rds  *redis.Client
	conf Conf
}

func New(c Conf) *RedisClient {
	options := &redis.Options{
		Addr: fmt.Sprintf("%s:%d", c.Host, c.Port),
		DB:   c.Db,
	}
	if c.Auth != "" {
		options.Password = c.Auth
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
	rds := redis.NewClient(options)
	return &RedisClient{
		rds:  rds,
		conf: c,
	}
}

func (c *RedisClient) Client() *redis.Client {
	return c.rds
}

func (c *RedisClient) Select(ctx context.Context, db int) *redis.Client {
	pipe := c.rds.Pipeline()
	_ = pipe.Select(ctx, db)
	if _, err := pipe.Exec(ctx); err != nil {
		panic(err)
	}
	return c.rds
}
