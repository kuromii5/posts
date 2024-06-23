package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/kuromii5/posts/internal/db"
	"github.com/kuromii5/posts/internal/models"
	"gorm.io/gorm"
)

func (d *PostgresDB) SaveUser(ctx context.Context, user *models.User) error {
	const f = "postgres.SaveUser"

	if err := d.DB.WithContext(ctx).Create(&user).Error; err != nil {
		return fmt.Errorf("%s:%w", f, err)
	}

	return nil
}

func (d *PostgresDB) UserByID(ctx context.Context, id uint64) (*models.User, error) {
	const f = "postgres.UserByID"

	var user models.User
	if err := d.DB.WithContext(ctx).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%s:%w", f, db.ErrNotFound)
		}

		return nil, fmt.Errorf("%s:%w", f, err)
	}

	return &user, nil
}
