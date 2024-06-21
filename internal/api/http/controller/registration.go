package controller

import (
	"encoding/json"
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
	cmd := inprocess.RegisterNewUserCommand{}

	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		logger.FromContext(r.Context()).Error(
			"Server process request with error.", "err", err, "body", r.Body,
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = c.registrationAPI.RegisterNewUser(r.Context(), cmd)

	if err != nil {
		logger.FromContext(r.Context()).Error(
			"Error occurred when server tried to register new user.", "err", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	logger.FromContext(r.Context()).Info("Successfully registered new user")
	w.WriteHeader(http.StatusCreated)
}

func (c *registrationHTTPController) RegistrationConfirm(
	w http.ResponseWriter,
	r *http.Request,
	registrationID server.RegistrationUuid,
) {
	logger.FromContext(r.Context()).Info("registration confirm.", "regId", registrationID)
	cmd := inprocess.ConfirmRegistrationCommand{RegistrationID: registrationID}

	_, err := c.registrationAPI.ConfirmRegistration(r.Context(), cmd)
	if err != nil {
		logger.FromContext(r.Context()).Error(
			"Error occurred when server tried to confirm registration.", "err", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}
