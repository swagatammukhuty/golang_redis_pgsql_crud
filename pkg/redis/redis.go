package redis

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	ctx     = context.Background()
	RedisDB *redis.Client
)

func InitRedis() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	_, err := RedisDB.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect with Redis %v", err)
	}
}
