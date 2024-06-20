package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisDB struct {
	client *redis.Client
}

func New(redisUrl string) (*RedisDB, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: "",
		DB:       0,
	})

	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	return &RedisDB{client: rdb}, nil
}

func (r *RedisDB) Close() error {
	if err := r.client.Close(); err != nil {
		return fmt.Errorf("close db error: %s", err)
	}
	return nil
}
