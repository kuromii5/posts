package redis

import (
	"context"

	"github.com/kuromii5/posts/internal/models"
)

func (d *RedisDB) SavePost(ctx context.Context, post *models.Post) error {
	panic("not implemented")
}

func (d *RedisDB) PostByID(ctx context.Context, id uint64) (*models.Post, error) {
	panic("not implemented")
}

func (d *RedisDB) Posts(ctx context.Context) ([]*models.Post, error) {
	panic("not implemented")
}
