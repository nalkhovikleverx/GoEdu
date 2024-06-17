package controller

import (
	"net/http"

	"go.opentelemetry.io/otel/trace"

	"GoEdu/internal/api/http/server"
	"GoEdu/internal/pkg/logger"
	"GoEdu/internal/registration/api/inprocess"
)

type registrationHTTPController struct {
	registrationAPI inprocess.RegistrationModuleFacade
	tracer          trace.Tracer
}

func newRegistrationHTTPController(
	t trace.Tracer,
	rAPI inprocess.RegistrationModuleFacade,
) *registrationHTTPController {
	return &registrationHTTPController{tracer: t, registrationAPI: rAPI}
}

func (c *registrationHTTPController) RegistrationNew(w http.ResponseWriter, r *http.Request) {
	logger.FromContext(r.Context()).Info("registration new")
	_, _ = c.registrationAPI.RegisterNewUser(r.Context(), inprocess.RegisterNewUserCommand{})
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
