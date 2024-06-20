package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/kuromii5/posts/internal/models"
	"gorm.io/gorm"
)

func (d *PostgresDB) SaveUser(ctx context.Context, username, email string, passwordHash []byte) error {
	const f = "postgres.SaveUser"

	// Check if email already exists
	var user models.User
	err := d.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return fmt.Errorf("%s: failed to check for existing user: %w", f, err)
	}
	if err == nil {
		return fmt.Errorf("%s: user already exists", f)
	}

	user = models.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(passwordHash),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if err := d.db.WithContext(ctx).Create(&user).Error; err != nil {
		return fmt.Errorf("%s: failed to create user: %w", f, err)
	}

	return nil
}

func (d *PostgresDB) User(ctx context.Context, email string) (*models.User, error) {
	const f = "postgres.User"

	var user models.User
	err := d.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("%s: user not found: %w", f, err)
		}
		return nil, fmt.Errorf("%s: failed to fetch user: %w", f, err)
	}

	return &user, nil
}
