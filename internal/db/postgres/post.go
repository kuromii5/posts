package postgres

import (
	"context"
	"fmt"

	"github.com/kuromii5/posts/internal/models"
)

func (d *PostgresDB) SavePost(ctx context.Context, post *models.Post) error {
	const f = "postgres.SavePost"

	if err := d.db.WithContext(ctx).Create(&post).Error; err != nil {
		return fmt.Errorf("%s: failed to create post: %w", f, err)
	}

	return nil
}

func (d *PostgresDB) PostByID(ctx context.Context, id uint64) (*models.Post, error) {
	const f = "postgres.PostByID"

	var post models.Post
	if err := d.db.WithContext(ctx).First(&post, id).Error; err != nil {
		return nil, fmt.Errorf("%s: failed to retrieve post: %w", f, err)
	}
	return &post, nil
}

func (d *PostgresDB) Posts(ctx context.Context) ([]*models.Post, error) {
	const f = "postgres.Posts"

	var posts []*models.Post
	if err := d.db.WithContext(ctx).Find(&posts).Error; err != nil {
		return nil, fmt.Errorf("%s: failed to retrieve posts: %w", f, err)
	}

	return posts, nil
}
