package storage

import (
	"fmt"
	"log"

	"github.com/kuromii5/posts/internal/config"
	"github.com/kuromii5/posts/internal/db"
	"github.com/kuromii5/posts/internal/db/postgres"
	"github.com/kuromii5/posts/internal/db/redis"
)

func New(config *config.Config) (db.DB, error) {
	const f = "db/db.New"
	var db db.DB
	var err error

	// choose database instance
	switch config.Storage {
	case "postgres":
		db, err = postgres.New(config.Postgres.URL)
	case "redis":
		db, err = redis.New(config.Redis.URL)
	default:
		return nil, fmt.Errorf("unsupported storage type: %s", config.Storage)
	}
	if err != nil {
		return nil, fmt.Errorf("%s:%w", f, err)
	}

	log.Println("Database connection established, storage:", config.Storage)

	return db, nil
}
