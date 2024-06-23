package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/kuromii5/posts/internal/db"
	l "github.com/kuromii5/posts/internal/lib/logger"
	"github.com/kuromii5/posts/internal/models"
)

type PostManager interface {
	SavePost(ctx context.Context, post *models.Post) error
	PostByID(ctx context.Context, id uint64) (*models.Post, error)
	Posts(ctx context.Context) ([]*models.Post, error)
}

type PostReg struct {
	Title   string `validate:"required,max=200"`
	Content string `validate:"required,max=10000"`
}

func (s *Service) CreatePost(ctx context.Context, title, content string, userID uint64, commentsEnabled bool) (*models.Post, error) {
	const f = "service.CreatePost"

	log := s.log.With(slog.String("func", f))
	log.Info("creating new post")

	// Validate input
	if err := s.validatePostReg(title, content); err != nil {
		log.Error("invalid title or content", l.Err(err))

		return nil, fmt.Errorf("%s:%w", f, err)
	}

	post := &models.Post{
		Title:           title,
		Content:         content,
		UserID:          userID,
		CommentsEnabled: commentsEnabled,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	err := s.PostService.SavePost(ctx, post)
	if err != nil {
		log.Error("failed to save post", l.Err(err))

		return nil, fmt.Errorf("%s:%w", f, err)
	}

	log.Info("post created successfully", slog.Uint64("post id", post.ID))

	return post, nil
}

func (s *Service) validatePostReg(title, content string) error {
	validate := validator.New()
	v := &PostReg{Title: title, Content: content}
	if err := validate.Struct(v); err != nil {
		return err
	}

	return nil
}

func (s *Service) PostByID(ctx context.Context, id uint64) (*models.Post, error) {
	const f = "service.PostByID"

	log := s.log.With(slog.String("func", f), slog.Uint64("post id", id))
	log.Info("retrieving post by post id")

	post, err := s.PostService.PostByID(ctx, id)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			log.Warn("post not found", l.Err(err))

			return nil, fmt.Errorf("%s:%w", f, ErrNotFound)
		}

		return nil, fmt.Errorf("%s:%w", f, err)
	}

	log.Info("post retrieved successfully")

	return post, nil
}

func (s *Service) Posts(ctx context.Context) ([]*models.Post, error) {
	const f = "service.Posts"

	log := s.log.With(slog.String("func", f))
	log.Info("retrieving all posts")

	posts, err := s.PostService.Posts(ctx)
	if err != nil {
		log.Error("failed to retrieve posts", l.Err(err))

		return nil, fmt.Errorf("%s:%w", f, err)
	}

	log.Info("all posts retrieved successfully")

	return posts, nil
}
