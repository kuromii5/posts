package service

import (
	"errors"
	"log/slog"
)

var ErrNotFound = errors.New("not found")

type Service struct {
	UserService    UserManager
	PostService    PostManager
	CommentService CommentManager
	PubSubService  *SubscriptionManager
	log            *slog.Logger
}

func New(
	userService UserManager,
	postService PostManager,
	commentService CommentManager,
	log *slog.Logger,
) *Service {
	pubSubService := NewSubscriptionManager(log)
	return &Service{
		UserService:    userService,
		PostService:    postService,
		CommentService: commentService,
		PubSubService:  pubSubService,
		log:            log,
	}
}
