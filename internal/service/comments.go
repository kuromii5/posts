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

type CommentManager interface {
	SaveComment(ctx context.Context, comment *models.Comment) error
	CommentByID(ctx context.Context, commID uint64) (*models.Comment, error)
	CommentsByPostID(ctx context.Context, postID uint64, limit, offset int) ([]*models.Comment, error)
	RepliesByCommentID(ctx context.Context, commID uint64, limit, offset int) ([]*models.Comment, error)
}

type CommentReg struct {
	Content string `validate:"required,max=2000"`
}

func (s *Service) CreateComment(
	ctx context.Context,
	userID, postID uint64,
	parentCommentID *uint64,
	content string,
) (*models.Comment, error) {
	const f = "service.CreateComment"

	log := s.log.With(slog.String("func", f))
	log.Info("creating new comment")

	// Validate input
	if err := s.validateCommentReg(content); err != nil {
		log.Error("invalid comment content", l.Err(err))

		return nil, fmt.Errorf("input validation error: %w", err)
	}

	// Check if parentCommentID is provided
	if parentCommentID != nil {
		// Retrieve parent comment to verify postID
		parentComment, err := s.CommentService.CommentByID(ctx, *parentCommentID)
		if err != nil {
			log.Error("failed to retrieve parent comment", l.Err(err))

			return nil, fmt.Errorf("failed to retrieve parent comment: %w", err)
		}

		// Check if postID of parent comment matches postID provided
		if parentComment.PostID != postID {
			errMsg := fmt.Sprintf("postID of parent comment (%d) does not match postID provided (%d)", parentComment.PostID, postID)
			log.Error(errMsg)

			return nil, errors.New(errMsg)
		}
	}

	comment := &models.Comment{
		PostID:          postID,
		UserID:          userID,
		ParentCommentID: parentCommentID,
		Content:         content,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	err := s.CommentService.SaveComment(ctx, comment)
	if err != nil {
		log.Error("failed to save comment", l.Err(err))

		return nil, fmt.Errorf("failed to save comment: %w", err)
	}

	log.Info("comment created successfully", slog.Uint64("comment id", comment.ID))

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
	const f = "service.CommentByID"

	log := s.log.With(slog.String("func", f), slog.Uint64("comment id", commID))
	log.Info("retrieving comment by comment id")

	comment, err := s.CommentService.CommentByID(ctx, commID)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			log.Warn("comment not found", l.Err(err))

			return nil, fmt.Errorf("%s:%w", f, ErrNotFound)
		}

		return nil, fmt.Errorf("%s:%w", f, err)
	}

	log.Info("comment retrieved successfully")

	return comment, nil
}

func (s *Service) CommentsByPostID(ctx context.Context, postID uint64, limit, offset *int) ([]*models.Comment, error) {
	const f = "service.CommentsByPostID"

	log := s.log.With(slog.String("func", f), slog.Uint64("post id", postID))
	log.Info("retrieving comments by post ID")

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

	comments, err := s.CommentService.CommentsByPostID(ctx, postID, *limit, *offset)
	if err != nil {
		log.Error("failed to retrieve comments", l.Err(err))

		return nil, fmt.Errorf("%s: %w", f, err)
	}

	log.Info("all comments retrieved successfully")

	return comments, nil
}

func (s *Service) RepliesByCommentID(ctx context.Context, commID uint64, limit, offset *int) ([]*models.Comment, error) {
	const f = "service.RepliesByCommentID"

	log := s.log.With(slog.String("func", f), slog.Uint64("comment id", commID))
	log.Info("retrieving replies by comment ID")

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

	replies, err := s.CommentService.RepliesByCommentID(ctx, commID, *limit, *offset)
	if err != nil {
		log.Error("failed to retrieve replies", l.Err(err))

		return nil, fmt.Errorf("%s: %w", f, err)
	}

	log.Info("all replies retrieved successfully")

	return replies, nil
}
