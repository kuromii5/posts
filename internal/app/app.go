package app

import (
	"log/slog"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/kuromii5/posts/internal/app/gqlserver"
	"github.com/kuromii5/posts/internal/db"
	"github.com/kuromii5/posts/internal/service"
)

type App struct {
	Server *gqlserver.GQLServer
}

func New(
	port int,
	log *slog.Logger,
	dbUrl string,
	dbType string,
	secret string,
	expires time.Duration,
) *App {
	// init db
	db, err := db.New(dbUrl, dbType)
	if err != nil {
		panic("failed to init database")
	}

	validator := validator.New()
	service := service.New(db, db, db, validator, log, secret, expires)
	server := gqlserver.New(log, port, service)

	return &App{Server: server}
}
