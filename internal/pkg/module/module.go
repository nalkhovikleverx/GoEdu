package module

import (
	"context"
	"log/slog"
	"net/http"

	"go.opentelemetry.io/otel/trace"

	"GoEdu/internal/pkg/config"
	"GoEdu/internal/registration/api/inprocess"
)

// Dependencies is the interface that holds all infrastructural dependencies that a module may need.
type Dependencies interface {
	Logger() *slog.Logger
	Config() *config.Config
	HTTP() *http.ServeMux
	TraceProvider() trace.TracerProvider
	RegistrationAPI() inprocess.RegistrationModuleFacade
	SetRegistrationAPI(inprocess.RegistrationModuleFacade)
}

// Module is the interface that every module must implement.
type Module interface {
	Init(context.Context, Dependencies) error
}
