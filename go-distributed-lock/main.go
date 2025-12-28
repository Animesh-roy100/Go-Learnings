package main

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

const LOCK_KEY = "resource_lock"
const EXPIRE_TIME = 10 * time.Second

// Acquire Lock
func acquireLock(redisClient *redis.Client, lockKey string, expireTime time.Duration) bool {
	result, err := redisClient.SetNX(ctx, lockKey, "locked", expireTime).Result()
	if err != nil {
		log.Fatalf("Error acquiring lock: %v", err)
	}
	return result
}

// Release Lock
func releaseLock(redisClient *redis.Client, lockKey string) {
	err := redisClient.Del(ctx, lockKey).Err()
	if err != nil {
		log.Fatalf("Error releasing lock: %v", err)
	}
}

// Usage
func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	if acquireLock(redisClient, LOCK_KEY, EXPIRE_TIME) {
		log.Println("Lock acquired")
		time.Sleep(5 * time.Second) // Simulating work
		releaseLock(redisClient, LOCK_KEY)
		log.Println("Lock released")
	} else {
		log.Println("Failed to acquire lock")
	}
}
