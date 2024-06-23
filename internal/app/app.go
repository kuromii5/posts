package app

import (
	"log/slog"

	"github.com/kuromii5/posts/internal/app/gqlserver"
	"github.com/kuromii5/posts/internal/app/storage"
	"github.com/kuromii5/posts/internal/config"
	"github.com/kuromii5/posts/internal/service"
)

type App struct {
	Server *gqlserver.GQLServer
}

func New(log *slog.Logger, config *config.Config) *App {
	// init db
	db, err := storage.New(config)
	if err != nil {
		panic("failed to init database")
	}

	// init service
	service := service.New(db, db, db, log)

	// init server
	server := gqlserver.New(log, config.Port, service)

	return &App{Server: server}
}
