package logger

import (
	"log/slog"
	"os"
)

var (
	LvlDebug = "debug"
	LvlInfo  = "info"
)

func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case LvlDebug:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case LvlInfo:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
