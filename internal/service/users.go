package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/kuromii5/posts/internal/db"

	jwtauth "github.com/kuromii5/posts/internal/lib/jwtAuth"
	l "github.com/kuromii5/posts/internal/lib/logger"
	"github.com/kuromii5/posts/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserSaver interface {
	SaveUser(ctx context.Context, username, email string, passwordHash []byte) error
}

type UserFetcher interface {
	User(ctx context.Context, email string) (*models.User, error)
}

type UserManager interface {
	UserSaver
	UserFetcher
}

func (s *Service) CreateUser(ctx context.Context, username string, email string, password string) (*models.User, error) {
	const f = "service.CreateUser"

	log := s.log.With(slog.String("func", f))
	log.Info("Registering new user")

	// Validate input
	if err := s.validateUserReg(username, email, password); err != nil {
		return nil, fmt.Errorf("input validation error: %w", err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to hash password", l.Err(err))
		return nil, fmt.Errorf("%s:%w", f, err)
	}

	if err = s.userService.SaveUser(ctx, username, email, hash); err != nil {
		log.Error("failed to save user", l.Err(err))
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	s.log.Info("user registered successfully", slog.String("username", username), slog.String("email", email))

	return &models.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hash),
	}, nil
}

func (s *Service) LoginUser(ctx context.Context, email, password string) (string, error) {
	const f = "service.LoginUser"

	log := s.log.With(slog.String("func", f))
	log.Info("user trying to log in")

	// get user from db
	user, err := s.userService.User(ctx, email)
	if err != nil {
		if errors.Is(err, db.ErrUserNotFound) {
			s.log.Warn("user not found", l.Err(err))
			return "", fmt.Errorf("%s:%w", f, ErrInvalidCredentials)
		}

		s.log.Error("failed to fetch user", l.Err(err))
		return "", fmt.Errorf("%s:%w", f, err)
	}

	// compare passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		s.log.Warn("invalid credentials", l.Err(err))
		return "", fmt.Errorf("%s:%w", f, ErrInvalidCredentials)
	}

	// generate jwt token
	token, err := jwtauth.GenerateJWTToken(user.ID, user.Email, s.secret, s.expires)
	if err != nil {
		s.log.Error("failed to generate jwt", l.Err(err))
		return "", fmt.Errorf("%s:%w", f, err)
	}

	log.Info("user logged in successfully")

	return token, nil
}
