package service

import (
	"context"

	"github.com/kuromii5/posts/internal/models"
)

type PostCreator interface {
	CreatePost(ctx context.Context, post *models.Post) error
}

type PostFetcher interface {
	Post(ctx context.Context, postID string) (*models.Post, error)
	Posts(ctx context.Context) ([]*models.Post, error)
}

type PostManager interface {
	PostCreator
	PostFetcher
}
