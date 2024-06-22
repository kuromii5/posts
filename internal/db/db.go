package db

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/kuromii5/posts/internal/config"
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
	RepliesByCommentID(ctx context.Context, commID uint64, limit, offset int) ([]*models.Comment, error)

	// Close function
	Close() error
}

func New(config *config.Config) (DB, error) {
	const f = "db/db.New"
	var db DB
	var err error

	switch config.Storage {
	case "postgres":
		db, err = postgres.New(config.Postgres.URL)
	case "redis":
		db, err = redis.New(config.Redis.URL)
	default:
		return nil, fmt.Errorf("unsupported storage type: %s", config.Storage)
	}
	if err != nil {
		return nil, fmt.Errorf("%s:%w", f, err)
	}

	log.Println("Database connection established, storage:", config.Storage)

	return db, nil
}
