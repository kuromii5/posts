package service

import (
	"errors"
	"log/slog"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserExists         = errors.New("user already exists")
)

type Service struct {
	userService    UserManager
	postService    PostManager
	commentService CommentManager
	log            *slog.Logger
}

func New(
	userService UserManager,
	postService PostManager,
	commentService CommentManager,
	log *slog.Logger,
) *Service {
	return &Service{
		userService:    userService,
		postService:    postService,
		commentService: commentService,
		log:            log,
	}
}
