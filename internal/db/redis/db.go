package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisDB struct {
	Client *redis.Client
}

func New(redisUrl string) (*RedisDB, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: "",
		DB:       0,
	})

	// check connection
	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	return &RedisDB{Client: rdb}, nil
}

func (r *RedisDB) Close() error {
	if err := r.Client.Close(); err != nil {
		return fmt.Errorf("close db error: %s", err)
	}
	return nil
}
