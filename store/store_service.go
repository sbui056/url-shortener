package store

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

/*** Store Service Setup ***/

// struct to wrap Redis client
type StorageService struct {
	redisClient *redis.Client
}

// declaration of the storage service and redis context
var (
	storeService = &StorageService{}
	ctx = context.Background()
)

const CacheDuration = 6 * time.Hour

// initialize store service and store pointer
func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
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

/*** Store API Design and Implementation ***/
func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key URL | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}

func RetrieveInitialUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitalUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}