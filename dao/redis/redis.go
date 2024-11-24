package redis

import (
	"fmt"
	"web-app/settings"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func Init(cfg *settings.RedisConfig) error {
	// 初始化redis连接
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB: cfg.Db,
		PoolSize: cfg.PoolSize,
	})

	_, err := rdb.Ping(rdb.Context()).Result()
	return err
}

func Close() {
	_ = rdb.Close()
}