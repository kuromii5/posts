package service

import (
	"errors"
	"log/slog"
	"time"

	"github.com/go-playground/validator/v10"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserExists         = errors.New("user already exists")
)

type Service struct {
	userService    UserManager
	postService    PostManager
	commentService CommentManager
	v              *validator.Validate
	log            *slog.Logger
	secret         string
	expires        time.Duration
}

func New(
	userService UserManager,
	postService PostManager,
	commentService CommentManager,
	v *validator.Validate,
	log *slog.Logger,
	secret string,
	expires time.Duration,
) *Service {
	return &Service{
		userService:    userService,
		postService:    postService,
		commentService: commentService,
		v:              v,
		log:            log,
		secret:         secret,
		expires:        expires,
	}
}
