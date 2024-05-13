package module

import (
	"log/slog"
	"net/http"

	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"

	"GoEdu/internal/pkg/config"
	"GoEdu/internal/pkg/logger"
)

type dependency struct {
	config *config.Config
	logger *slog.Logger
	mux    *http.ServeMux
	tp     trace.TracerProvider
}

func BuildDefaultDependencies(cfg *config.Config) (Dependencies, error) {
	d := &dependency{config: cfg}
	d.initLogger()
	d.initHTTPMux()
	d.initTraceProvider()
	return d, nil
}

func (s *dependency) Logger() *slog.Logger {
	return s.logger
}

func (s *dependency) Config() *config.Config {
	return s.config
}

func (s *dependency) HTTP() *http.ServeMux {
	return s.mux
}

func (s *dependency) TraceProvider() trace.TracerProvider {
	return s.tp
}

func (s *dependency) initLogger() {
	var lvl slog.Level
	switch s.config.LogLevel {
	case "DEBUG":
		lvl = slog.LevelDebug
	case "ERROR":
		lvl = slog.LevelError
	case "WARN":
		lvl = slog.LevelWarn
	default:
		lvl = slog.LevelInfo
	}
	s.logger = logger.New(logger.Config{
		Level: lvl,
	})
}

func (s *dependency) initHTTPMux() {
	s.mux = http.DefaultServeMux
}

func (s *dependency) initTraceProvider() {
	s.tp = noop.NewTracerProvider()
}
