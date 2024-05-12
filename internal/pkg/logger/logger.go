// package logger provides a better defaults, configuration wrappers and utility functions for slog package.
package logger

import (
	"log/slog"
	"os"
)

type Config struct {
	Level slog.Level
}

func New(cfg Config) *slog.Logger {
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     cfg.Level,
	}))
	slog.SetDefault(l)
	return l
}
