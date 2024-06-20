package logger

import (
	"os"

	"log/slog"

	offlog "github.com/kuromii5/posts/internal/lib/logger/off"
	prettylog "github.com/kuromii5/posts/internal/lib/logger/pretty"
)

func New(env string) *slog.Logger {
	switch env {
	case "local":
		return prettylog.New(os.Stdout, slog.LevelDebug)
	case "dev":
		return prettylog.New(os.Stdout, slog.LevelDebug)
	case "prod":
		return prettylog.New(os.Stdout, slog.LevelInfo)
	default:
		return offlog.New()
	}
}

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
