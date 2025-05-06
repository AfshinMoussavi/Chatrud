package redisPkg

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type IRedis interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
}
