package main

import (
	"log/slog"

	"github.com/kuromii5/posts/internal/app"
	"github.com/kuromii5/posts/internal/config"
	l "github.com/kuromii5/posts/internal/lib/logger"
)

func main() {
	// Load config
	cfg := config.MustLoad()

	// init logger
	log := l.New(cfg.Env)
	slog.SetDefault(log)

	// init app
	app := app.New(cfg.Port, log, cfg.DBUrl, cfg.Storage)

	// run app
	app.Server.MustRun()
}
