package utils

import (
	"context"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/redis/go-redis/v9"
)

type CacheManager struct {
	inMemoryCache *cache.Cache
	redisClient   *redis.Client
	ctx           context.Context
}

func NewCacheManager() *CacheManager {
	return &CacheManager{
		inMemoryCache: cache.New(5*time.Minute, 10*time.Minute),
		redisClient: redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		}),
		ctx: context.Background(),
	}
}

func (c *CacheManager) Get(key string) (string, bool) {
	// Check in-memory cache
	if val, found := c.inMemoryCache.Get(key); found {
		return val.(string), true
	}

	// Check Redis
	val, err := c.redisClient.Get(c.ctx, key).Result()
	if err == nil {
		c.inMemoryCache.Set(key, val, cache.DefaultExpiration)
		return val, true
	}

	return "", false
}

func (c *CacheManager) Set(key string, value string, expiration time.Duration) {
	c.inMemoryCache.Set(key, value, expiration)
	c.redisClient.Set(c.ctx, key, value, expiration)
}

func (c *CacheManager) Delete(key string) {
	c.inMemoryCache.Delete(key)
	c.redisClient.Del(c.ctx, key)
}
