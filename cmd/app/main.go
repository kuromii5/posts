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

	log.Debug("config", slog.Any("port", cfg.Port), slog.Any("postgres url", cfg.Postgres.URL), slog.Any("redis url", cfg.Redis.URL), slog.Any("storage type", cfg.Storage))

	// init app
	app := app.New(log, cfg)

	// run app
	app.Server.MustRun()
}
