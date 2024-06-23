package db

import (
	"context"
	"errors"

	"github.com/kuromii5/posts/internal/models"
)

var ErrNotFound = errors.New("not found")

type DB interface {
	// Save user in database
	SaveUser(ctx context.Context, user *models.User) error

	// Get user from database by ID
	UserByID(ctx context.Context, id uint64) (*models.User, error)

	// Save post in database
	SavePost(ctx context.Context, post *models.Post) error

	// Get post from database by ID
	PostByID(ctx context.Context, id uint64) (*models.Post, error)

	// Get all created posts
	Posts(ctx context.Context) ([]*models.Post, error)

	// Save comment in database
	SaveComment(ctx context.Context, comment *models.Comment) error

	// Get comment by ID
	CommentByID(ctx context.Context, commID uint64) (*models.Comment, error)

	// Get all comments on post with given ID (pagination included)
	CommentsByPostID(ctx context.Context, postID uint64, limit, offset int) ([]*models.Comment, error)

	// Get all replies on the comment with given ID (pagination included)
	RepliesByCommentID(ctx context.Context, commID uint64, limit, offset int) ([]*models.Comment, error)

	// Close db connection function
	Close() error
}
