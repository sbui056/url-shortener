package store

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// struct to wrap Redis client
type StoreService struct {
	client *redis.Client
}

// declaration of the storage service and redis context
var (
	storeService = &StorageService{}
	ctx = context.Background()
)

const CacheExpiration = 6 * time.Hour

// initialize store service and store pointer
func InitializeStore() *StorageService {
	redisClient := redi.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}