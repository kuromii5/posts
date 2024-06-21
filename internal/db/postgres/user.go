package postgres

import (
	"context"
	"fmt"

	"github.com/kuromii5/posts/internal/models"
)

func (d *PostgresDB) SaveUser(ctx context.Context, user *models.User) error {
	const f = "postgres.SaveUser"

	if err := d.db.WithContext(ctx).Create(&user).Error; err != nil {
		return fmt.Errorf("%s: failed to create user: %w", f, err)
	}
	return nil
}

func (d *PostgresDB) UserByID(ctx context.Context, id uint64) (*models.User, error) {
	const f = "postgres.UserByID"

	var user models.User
	if err := d.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, fmt.Errorf("%s: failed to retrieve user: %w", f, err)
	}
	return &user, nil
}
