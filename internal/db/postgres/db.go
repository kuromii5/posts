package postgres

import (
	"fmt"
	"log"
	"time"

	"github.com/kuromii5/posts/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	db *gorm.DB
}

func New(dbUrl string) (*PostgresDB, error) {
	db, err := gorm.Open(postgres.Open(dbUrl))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	log.Println("Database connection established")

	if err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return &PostgresDB{db: db}, nil
}

func (d *PostgresDB) Close() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}
	return sqlDB.Close()
}
