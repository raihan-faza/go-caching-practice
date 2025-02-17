package main

import (
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/redis/go-redis"
)

func init() {
	inMemCache := cache.New(10*time.Minute, 15*time.Minute)
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Update if Redis is running elsewhere
	})
}
