// package logger provides a better defaults, configuration wrappers and utility functions for slog package.
package logger

import (
	"log/slog"
	"os"
)

type VCSConfig struct {
	Revision  string
	Tag       string
	BuildTime string
}

type Config struct {
	Level   slog.Level
	Service string
	VCS     VCSConfig
}

func New(cfg Config) *slog.Logger {
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     cfg.Level,
	})).With(
		"vcs.revision", cfg.VCS.Revision,
		"vcs.tag", cfg.VCS.Tag,
		"vcs.time", cfg.VCS.BuildTime,
		"service", cfg.Service,
	)
	slog.SetDefault(l)
	return l
}
