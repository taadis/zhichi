package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var _ Cache = (*RedisCache)(nil)

type RedisCache struct {
	conn redis.UniversalClient
}

func NewRedisCache(conn redis.UniversalClient) *RedisCache {
	c := new(RedisCache)
	c.conn = conn
	return c
}

func (c *RedisCache) Get(ctx context.Context, key string) interface{} {
	ret, err := c.conn.Get(ctx, key).Result()
	if err != nil {
		return nil
	}
	return ret
}

func (c *RedisCache) Set(ctx context.Context, key string, v interface{}, d time.Duration) error {
	return c.conn.SetEx(ctx, key, v, d).Err()
}

func (c *RedisCache) Has(ctx context.Context, key string) bool {
	ret, err := c.conn.Exists(ctx, key).Result()
	if err != nil {
		return false
	}
	return ret > 0
}

func (c *RedisCache) Del(ctx context.Context, key string) error {
	return c.conn.Del(ctx, key).Err()
}
