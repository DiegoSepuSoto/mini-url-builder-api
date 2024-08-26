package shared

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

const defaultExpirationTime = 24 * time.Hour

type RedisClient interface {
	Set(ctx context.Context, key, value string) RedisStringResult
}

type RedisStringResult interface {
	Result() (string, error)
}

type redisClient struct {
	redisClient *redis.Client
}

func (c *redisClient) Set(ctx context.Context, key, value string) RedisStringResult {
	return c.redisClient.Set(ctx, key, value, defaultExpirationTime)
}

func CreateRedisClient() *redisClient {
	return &redisClient{
		redisClient: redis.NewClient(&redis.Options{
			Addr: os.Getenv("REDIS_HOST"),
		}),
	}
}
