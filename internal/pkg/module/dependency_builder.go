package module

import (
	"log/slog"
	"net/http"

	"GoEdu/internal/pkg/config"
	"GoEdu/internal/pkg/logger"
	"GoEdu/internal/registration/api/inprocess"

	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

type dependency struct {
	reg    inprocess.RegistrationModuleFacade
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

func (s *dependency) RegistrationAPI() inprocess.RegistrationModuleFacade {
	return s.reg
}

func (s *dependency) SetRegistrationAPI(f inprocess.RegistrationModuleFacade) {
	s.reg = f
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
		Service: s.config.ServiceName,
		Level:   lvl,
		VCS: logger.VCSConfig{
			Revision:  s.config.VCS.Revision,
			Tag:       s.config.VCS.Tag,
			BuildTime: s.config.VCS.Time,
		},
	})
}

func (s *dependency) initHTTPMux() {
	s.mux = http.DefaultServeMux
}

func (s *dependency) initTraceProvider() {
	s.tp = noop.NewTracerProvider()
}
