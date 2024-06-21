package redis

import (
	"context"

	"github.com/kuromii5/posts/internal/models"
)

func (d *RedisDB) SaveComment(ctx context.Context, comment *models.Comment) error {
	panic("not implemented")
}

func (d *RedisDB) CommentsByPostID(ctx context.Context, postID uint64, limit, offset int) ([]*models.Comment, error) {
	panic("not implemented")
}

func (d *RedisDB) CommentByID(ctx context.Context, commID uint64) (*models.Comment, error) {
	panic("not implemented")
}
