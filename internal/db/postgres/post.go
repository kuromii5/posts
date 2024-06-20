package postgres

import (
	"context"

	"github.com/kuromii5/posts/internal/models"
)

func (d *PostgresDB) CreatePost(ctx context.Context, post *models.Post) error {
	panic("not implemented")
}

func (d *PostgresDB) Post(ctx context.Context, postID string) (*models.Post, error) {
	panic("not implemented")
}

func (d *PostgresDB) Posts(ctx context.Context) ([]*models.Post, error) {
	panic("not implemented")
}
