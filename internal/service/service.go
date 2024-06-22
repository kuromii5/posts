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
	PubSubService  *SubscriptionManager
	log            *slog.Logger
}

func New(
	userService UserManager,
	postService PostManager,
	commentService CommentManager,
	log *slog.Logger,
) *Service {
	pubSubService := NewSubscriptionManager()
	return &Service{
		userService:    userService,
		postService:    postService,
		commentService: commentService,
		PubSubService:  pubSubService,
		log:            log,
	}
}
