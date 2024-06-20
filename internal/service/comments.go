package service

import (
	"context"

	"github.com/kuromii5/posts/internal/models"
)

type CommentCreator interface {
	CreateComment(ctx context.Context, comment *models.Comment) error
}

type CommentFetcher interface {
	CommentsByPost(ctx context.Context, postID string) ([]*models.Comment, error)
	CommentsByUser(ctx context.Context, userID string) ([]*models.Comment, error)
}

type CommentManager interface {
	CommentCreator
	CommentFetcher
}
