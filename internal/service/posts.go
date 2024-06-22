package service

import (
	"context"
	"log/slog"
	"time"

	l "github.com/kuromii5/posts/internal/lib/logger"
	"github.com/kuromii5/posts/internal/models"
)

type PostManager interface {
	SavePost(ctx context.Context, post *models.Post) error
	PostByID(ctx context.Context, id uint64) (*models.Post, error)
	Posts(ctx context.Context) ([]*models.Post, error)
}

func (s *Service) CreatePost(ctx context.Context, title, content string, userID uint64, commentsEnabled bool) (*models.Post, error) {
	const f = "service.CreatePost"

	log := s.log.With(slog.String("func", f))
	log.Info("Creating new post")

	post := &models.Post{
		Title:           title,
		Content:         content,
		UserID:          userID,
		CommentsEnabled: commentsEnabled,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	err := s.postService.SavePost(ctx, post)
	if err != nil {
		log.Error("failed to save post", l.Err(err))
		return nil, err
	}

	log.Info("post created successfully", slog.Uint64("userID", userID))

	return post, nil
}

func (s *Service) PostByID(ctx context.Context, id uint64) (*models.Post, error) {
	return s.postService.PostByID(ctx, id)
}

func (s *Service) Posts(ctx context.Context) ([]*models.Post, error) {
	return s.postService.Posts(ctx)
}
