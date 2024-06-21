package app

import (
	"log/slog"

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
) *App {
	// init db
	db, err := db.New(dbUrl, dbType)
	if err != nil {
		panic("failed to init database")
	}

	// init service
	service := service.New(db, db, db, log)

	// init server
	server := gqlserver.New(log, port, service)

	return &App{Server: server}
}
