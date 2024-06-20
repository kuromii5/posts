package redis

import (
	"context"

	"github.com/kuromii5/posts/internal/models"
)

func (d *RedisDB) CreateComment(ctx context.Context, comment *models.Comment) error {
	panic("not implemented")
}

func (d *RedisDB) CommentsByPost(ctx context.Context, postID string) ([]*models.Comment, error) {
	panic("not implemented")
}

func (d *RedisDB) CommentsByUser(ctx context.Context, userID string) ([]*models.Comment, error) {
	panic("not implemented")
}
