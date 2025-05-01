package redisPkg

import (
	"Chat-Websocket/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var Rdb *redis.Client

func InitRedis(cfg *config.Config) error {
	address := fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)
	Rdb = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Db,
	})

	if _, err := Rdb.Ping(Ctx).Result(); err != nil {
		return fmt.Errorf("redis connection failed: %v", err)
	}
	return nil
}
