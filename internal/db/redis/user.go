package redis

import (
	"context"

	"github.com/kuromii5/posts/internal/models"
)

func (r *RedisDB) SaveUser(ctx context.Context, username, email string, passwordHash []byte) error {
	panic("not implemented")
}

func (r *RedisDB) User(ctx context.Context, email string) (*models.User, error) {
	panic("not implemented")
}
