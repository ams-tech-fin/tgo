package cache

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

func SetupRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       0,
	})
}

func PingRedis(ctx context.Context) (string, error) {
	return Rdb.Ping(ctx).Result()
}

func Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return Rdb.Set(ctx, key, value, expiration).Err()
}

func Get(ctx context.Context, key string) (string, error) {

	result, err := Rdb.Get(ctx, key).Result()

	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}

	return result, nil
}
