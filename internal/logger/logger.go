package logger

import (
	"log/slog"
	"os"
)

const (
	lvlLocal = "local"
	lvlDev   = "dev"
	lvlProd  = "prod"
)

func New(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case lvlLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case lvlDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case lvlProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
