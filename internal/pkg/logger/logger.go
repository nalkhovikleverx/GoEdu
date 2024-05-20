// package logger provides a better defaults, configuration wrappers and utility functions for slog package.
package logger

import (
	"log/slog"
	"os"
)

type Config struct {
	Level        slog.Level
	Service      string
	VCSRevision  string
	VCSTag       string
	VCSBuildTime string
}

func New(cfg Config) *slog.Logger {
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     cfg.Level,
	})).With(
		"vcs.revision", cfg.VCSRevision,
		"vcs.tag", cfg.VCSTag,
		"vcs.time", cfg.VCSBuildTime,
		"service", cfg.Service,
	)
	slog.SetDefault(l)
	return l
}
