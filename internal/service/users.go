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

type UserManager interface {
	SaveUser(ctx context.Context, user *models.User) error
	UserByID(ctx context.Context, id uint64) (*models.User, error)
}

type UserReg struct {
	username string `validate:"required,min=2"`
}

func (s *Service) CreateUser(ctx context.Context, username string) (*models.User, error) {
	const f = "service.CreateUser"

	log := s.log.With(slog.String("func", f), slog.String("username", username))
	log.Info("registering new user")

	// Validate input
	if err := s.validateUserReg(username); err != nil {
		log.Error("invalid username", l.Err(err))

		return nil, fmt.Errorf("%s:%w", f, err)
	}

	user := &models.User{
		Username:  username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := s.UserService.SaveUser(ctx, user)
	if err != nil {
		log.Error("failed to save user", l.Err(err))

		return nil, fmt.Errorf("%s:%w", f, err)
	}

	s.log.Info("user registered successfully", slog.Uint64("user id", user.ID))

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
	const f = "service.UserByID"

	log := s.log.With(slog.String("func", f), slog.Uint64("user id", id))
	log.Info("retrieving user by user id")

	user, err := s.UserService.UserByID(ctx, id)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			log.Warn("user not found", l.Err(err))

			return nil, fmt.Errorf("%s:%w", f, ErrNotFound)
		}

		return nil, fmt.Errorf("%s:%w", f, err)
	}

	log.Info("user retrieved successfully")

	return user, nil
}
