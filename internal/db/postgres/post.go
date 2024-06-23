package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/kuromii5/posts/internal/db"
	"github.com/kuromii5/posts/internal/models"
	"gorm.io/gorm"
)

func (d *PostgresDB) SavePost(ctx context.Context, post *models.Post) error {
	const f = "postgres.SavePost"

	if err := d.DB.WithContext(ctx).Create(&post).Error; err != nil {
		return fmt.Errorf("%s:%w", f, err)
	}

	return nil
}

func (d *PostgresDB) PostByID(ctx context.Context, id uint64) (*models.Post, error) {
	const f = "postgres.PostByID"

	var post models.Post
	if err := d.DB.WithContext(ctx).First(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%s:%w", f, db.ErrNotFound)
		}

		return nil, fmt.Errorf("%s:%w", f, err)
	}

	return &post, nil
}

func (d *PostgresDB) Posts(ctx context.Context) ([]*models.Post, error) {
	const f = "postgres.Posts"

	var posts []*models.Post
	if err := d.DB.WithContext(ctx).Find(&posts).Error; err != nil {
		return nil, fmt.Errorf("%s:%w", f, err)
	}

	return posts, nil
}
