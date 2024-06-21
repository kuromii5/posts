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

type UserManager interface {
	SaveUser(ctx context.Context, user *models.User) error
	UserByID(ctx context.Context, id uint64) (*models.User, error)
}

type UserReg struct {
	username string `validate:"required,min=2"`
}

func (s *Service) CreateUser(ctx context.Context, username string) (*models.User, error) {
	const f = "service.CreateUser"

	log := s.log.With(slog.String("func", f))
	log.Info("Registering new user")

	// Validate input
	if err := s.validateUserReg(username); err != nil {
		log.Error("invalid username", l.Err(err))
		return nil, fmt.Errorf("input validation error: %w", err)
	}

	user := &models.User{
		Username:  username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := s.userService.SaveUser(ctx, user)
	if err != nil {
		log.Error("failed to save user", l.Err(err))
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	s.log.Info("user registered successfully", slog.String("username", username))

	return user, nil
}

func (s *Service) validateUserReg(username string) error {
	validate := validator.New()
	v := &UserReg{username: username}
	if err := validate.Struct(v); err != nil {
		return err
	}
	return nil
}

func (s *Service) UserByID(ctx context.Context, id uint64) (*models.User, error) {
	return s.userService.UserByID(ctx, id)
}
