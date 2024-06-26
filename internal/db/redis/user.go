package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kuromii5/posts/internal/db"
	"github.com/kuromii5/posts/internal/models"
	"github.com/redis/go-redis/v9"
)

func (d *RedisDB) SaveUser(ctx context.Context, user *models.User) error {
	const f = "redis.SaveUser"

	// generate unique ID
	userID, err := d.Client.Incr(ctx, "user:id").Result()
	if err != nil {
		return fmt.Errorf("%s: failed to increment user ID: %w", f, err)
	}
	user.ID = uint64(userID)

	userData, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("%s:%w", f, err)
	}

	key := fmt.Sprintf("user:%d", user.ID)
	if err := d.Client.Set(ctx, key, userData, 0).Err(); err != nil {
		return fmt.Errorf("%s:%w", f, err)
	}

	return nil
}

func (r *RedisDB) UserByID(ctx context.Context, id uint64) (*models.User, error) {
	const f = "redis.UserByID"

	key := fmt.Sprintf("user:%d", id)

	var user models.User
	data, err := r.Client.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("%s:%w", f, db.ErrNotFound)
		}
		return nil, fmt.Errorf("%s:%v", f, err)
	}

	err = json.Unmarshal(data, &user)
	if err != nil {
		return nil, fmt.Errorf("%s:%v", f, err)
	}

	return &user, nil
}
