package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/kuromii5/posts/internal/db"
	"github.com/kuromii5/posts/internal/models"
	"gorm.io/gorm"
)

func (d *PostgresDB) SaveComment(ctx context.Context, comment *models.Comment) error {
	const f = "postgres.SaveComment"

	if err := d.DB.WithContext(ctx).Create(&comment).Error; err != nil {
		return fmt.Errorf("%s:%w", f, err)
	}

	return nil
}

func (d *PostgresDB) CommentByID(ctx context.Context, id uint64) (*models.Comment, error) {
	const f = "postgres.CommentByID"

	var comment models.Comment
	if err := d.DB.WithContext(ctx).First(&comment, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%s:%w", f, db.ErrNotFound)
		}

		return nil, fmt.Errorf("%s:%w", f, err)
	}

	return &comment, nil
}

func (d *PostgresDB) CommentsByPostID(ctx context.Context, postID uint64, limit, offset int) ([]*models.Comment, error) {
	const f = "postgres.CommentsByPostID"

	var comments []*models.Comment
	err := d.DB.WithContext(ctx).
		Where("post_id = ? AND parent_comment_id IS NULL", postID).
		Limit(limit).
		Offset(offset).
		Find(&comments).
		Error
	if err != nil {
		return nil, fmt.Errorf("%s:%w", f, err)
	}

	return comments, nil
}

func (d *PostgresDB) RepliesByCommentID(ctx context.Context, commID uint64, limit, offset int) ([]*models.Comment, error) {
	const f = "postgres.RepliesByCommentID"

	var comments []*models.Comment
	err := d.DB.WithContext(ctx).
		Where("parent_comment_id = ?", commID).
		Limit(limit).
		Offset(offset).
		Find(&comments).
		Error
	if err != nil {
		return nil, fmt.Errorf("%s:%w", f, err)
	}

	return comments, nil
}
