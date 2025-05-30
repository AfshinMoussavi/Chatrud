package redisPkg

import (
	"Chat-Websocket/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var Ctx = context.Background()
var Rdb *redis.Client

type RealRedisClient struct {
	Client *redis.Client
}

func InitRedis(cfg *config.Config) (IRedis, error) {
	address := fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Db,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, fmt.Errorf("redis connection failed: %v", err)
	}

	return &RealRedisClient{Client: client}, nil
}

func (r *RealRedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	return r.Client.Get(ctx, key)
}

func (r *RealRedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.Client.Set(ctx, key, value, expiration)
}

func (r *RealRedisClient) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	return r.Client.Del(ctx, keys...)
}
