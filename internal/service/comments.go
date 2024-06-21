package service

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/go-playground/validator/v10"
	l "github.com/kuromii5/posts/internal/lib/logger"
	"github.com/kuromii5/posts/internal/models"
)

type CommentManager interface {
	SaveComment(ctx context.Context, comment *models.Comment) error
	CommentByID(ctx context.Context, commID uint64) (*models.Comment, error)
	CommentsByPostID(ctx context.Context, postID uint64, limit, offset int) ([]*models.Comment, error)
}

type CommentReg struct {
	Content string `validate:"required,max=2000"`
}

func (s *Service) CreateComment(
	ctx context.Context,
	postID, userID uint64,
	parentCommentID *uint64,
	content string,
) (*models.Comment, error) {
	const f = "service.CreateComment"

	log := s.log.With(slog.String("func", f))
	log.Info("Creating new comment")

	// Validate input
	if err := s.validateCommentReg(content); err != nil {
		log.Error("invalid comment content", l.Err(err))
		return nil, fmt.Errorf("input validation error: %w", err)
	}

	comment := &models.Comment{
		PostID:          postID,
		UserID:          userID,
		ParentCommentID: parentCommentID,
		Content:         content,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	err := s.commentService.SaveComment(ctx, comment)
	if err != nil {
		log.Error("failed to save comment", l.Err(err))
		return nil, fmt.Errorf("failed to save comment: %w", err)
	}

	log.Info("comment created successfully", slog.Uint64("postID", postID), slog.Uint64("userID", userID))

	return comment, nil
}

func (s *Service) validateCommentReg(content string) error {
	validate := validator.New()
	v := &CommentReg{Content: content}
	if err := validate.Struct(v); err != nil {
		return err
	}
	return nil
}

func (s *Service) CommentByID(ctx context.Context, commID uint64) (*models.Comment, error) {
	return s.commentService.CommentByID(ctx, commID)
}

func (s *Service) CommentsByPostID(ctx context.Context, postID uint64, limit, offset *int) ([]*models.Comment, error) {
	// Use default limit if not provided
	if limit == nil {
		defaultLimit := 10
		limit = &defaultLimit
	}

	// Use default offset if not provided
	if offset == nil {
		defaultOffset := 0
		offset = &defaultOffset
	}

	return s.commentService.CommentsByPostID(ctx, postID, *limit, *offset)
}
