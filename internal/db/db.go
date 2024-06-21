package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/kuromii5/posts/internal/db/postgres"
	"github.com/kuromii5/posts/internal/db/redis"
	"github.com/kuromii5/posts/internal/models"
)

var (
	ErrUserExists   = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
)

type DB interface {
	// User management
	SaveUser(ctx context.Context, user *models.User) error
	UserByID(ctx context.Context, id uint64) (*models.User, error)

	// Post management
	SavePost(ctx context.Context, post *models.Post) error
	PostByID(ctx context.Context, id uint64) (*models.Post, error)
	Posts(ctx context.Context) ([]*models.Post, error)

	// Comment managemenet
	SaveComment(ctx context.Context, comment *models.Comment) error
	CommentByID(ctx context.Context, commID uint64) (*models.Comment, error)
	CommentsByPostID(ctx context.Context, postID uint64, limit, offset int) ([]*models.Comment, error)

	// Close function
	Close() error
}

func New(dbUrl string, dbType string) (DB, error) {
	const f = "db/db.New"
	var db DB
	var err error

	switch dbType {
	case "postgres":
		db, err = postgres.New(dbUrl)
	case "redis":
		db, err = redis.New(dbUrl)
	default:
		return nil, fmt.Errorf("unsupported storage type: %s", dbType)
	}
	if err != nil {
		return nil, fmt.Errorf("%s:%w", f, err)
	}

	return db, nil
}
