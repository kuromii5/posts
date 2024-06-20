package redis

import (
	"context"

	"github.com/kuromii5/posts/internal/models"
)

func (d *RedisDB) CreatePost(ctx context.Context, post *models.Post) error {
	panic("not implemented")
}

func (d *RedisDB) Post(ctx context.Context, postID string) (*models.Post, error) {
	panic("not implemented")
}

func (d *RedisDB) Posts(ctx context.Context) ([]*models.Post, error) {
	panic("not implemented")
}
