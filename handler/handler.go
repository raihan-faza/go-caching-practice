package handler

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

func getUserInfo(userID string) string {
	// 1. Check in-memory cache
	if val, found := inMemoryCache.Get(userID); found {
		fmt.Println("Cache Hit: In-memory")
		return val.(string)
	}

	// 2. Check Redis cache
	val, err := redisClient.Get(ctx, userID).Result()
	if err == nil {
		fmt.Println("Cache Hit: Redis")
		// Store in in-memory cache for faster access next time
		inMemoryCache.Set(userID, val, cache.DefaultExpiration)
		return val
	}

	// 3. Fetch from Database (Simulated)
	fmt.Println("Cache Miss: Fetching from DB")
	data := fetchFromDatabase(userID)

	// Store in both caches
	inMemoryCache.Set(userID, data, cache.DefaultExpiration)
	redisClient.Set(ctx, userID, data, 10*time.Minute)

	return data
}

// Simulated database fetch
func fetchFromDatabase(userID string) string {
	// Simulating a slow database query
	time.Sleep(2 * time.Second)
	return "UserData-" + userID
}
