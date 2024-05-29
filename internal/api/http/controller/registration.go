package controller

import (
	"net/http"

	"go.opentelemetry.io/otel/trace"

	"GoEdu/internal/api/http/server"
	"GoEdu/internal/pkg/logger"
)

type registrationHTTPController struct {
	tracer trace.Tracer
}

func newRegistrationHTTPController(t trace.Tracer) *registrationHTTPController {
	return &registrationHTTPController{tracer: t}
}

func (c *registrationHTTPController) RegistrationNew(w http.ResponseWriter, r *http.Request) {
	logger.FromContext(r.Context()).Info("registration new")
	w.WriteHeader(http.StatusNotImplemented)
}

func (c *registrationHTTPController) RegistrationConfirm(
	w http.ResponseWriter,
	r *http.Request,
	registrationID server.RegistrationUuid,
) {
	ctx := r.Context()
	l := logger.FromContext(ctx).With("registration_id", registrationID)
	l.Info("registration confirmed")
	w.WriteHeader(http.StatusNotImplemented)
}
