package postgres

import (
	"context"
	"fmt"

	"github.com/kuromii5/posts/internal/models"
)

func (d *PostgresDB) SaveComment(ctx context.Context, comment *models.Comment) error {
	const f = "postgres.SaveComment"

	if err := d.db.WithContext(ctx).Create(&comment).Error; err != nil {
		return fmt.Errorf("%s: failed to save comment: %w", f, err)
	}

	return nil
}

func (d *PostgresDB) CommentByID(ctx context.Context, id uint64) (*models.Comment, error) {
	const f = "postgres.CommentByID"

	var comment models.Comment
	if err := d.db.WithContext(ctx).First(&comment, id).Error; err != nil {
		return nil, fmt.Errorf("%s: failed to retrieve post: %w", f, err)
	}
	return &comment, nil
}

func (d *PostgresDB) CommentsByPostID(ctx context.Context, postID uint64, limit, offset int) ([]*models.Comment, error) {
	const f = "postgres.CommentsByPostID"

	var comments []*models.Comment
	err := d.db.WithContext(ctx).
		Where("post_id = ? AND parent_comment_id IS NULL", postID).
		Limit(limit).
		Offset(offset).
		Find(&comments).Error
	if err != nil {
		return nil, fmt.Errorf("%s: failed to retrieve comments: %w", f, err)
	}
	return comments, nil
}

func (d *PostgresDB) RepliesByCommentID(ctx context.Context, commID uint64, limit, offset int) ([]*models.Comment, error) {
	const f = "postgres.RepliesByCommentID"

	var comments []*models.Comment
	err := d.db.WithContext(ctx).Where("parent_comment_id = ?", commID).Limit(limit).Offset(offset).Find(&comments).Error
	if err != nil {
		return nil, fmt.Errorf("%s: failed to retrieve replies: %w", f, err)
	}
	return comments, nil
}
