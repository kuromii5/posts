package postgres

import (
	"fmt"

	"github.com/kuromii5/posts/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	DB *gorm.DB
}

func New(dbUrl string) (*PostgresDB, error) {
	db, err := gorm.Open(postgres.Open(dbUrl))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// automigrations
	if err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return &PostgresDB{DB: db}, nil
}

func (d *PostgresDB) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}
	return sqlDB.Close()
}
