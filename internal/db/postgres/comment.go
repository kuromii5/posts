package postgres

import (
	"context"

	"github.com/kuromii5/posts/internal/models"
)

func (d *PostgresDB) CreateComment(ctx context.Context, comment *models.Comment) error {
	panic("not implemented")
}

func (d *PostgresDB) CommentsByPost(ctx context.Context, postID string) ([]*models.Comment, error) {
	panic("not implemented")
}

func (d *PostgresDB) CommentsByUser(ctx context.Context, userID string) ([]*models.Comment, error) {
	panic("not implemented")
}
