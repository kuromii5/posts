package redis

import (
	"context"

	"github.com/kuromii5/posts/internal/models"
)

func (r *RedisDB) SaveUser(ctx context.Context, user *models.User) error {
	panic("not implemented")
}

func (r *RedisDB) UserByID(ctx context.Context, id uint64) (*models.User, error) {
	panic("not implemented")
}
