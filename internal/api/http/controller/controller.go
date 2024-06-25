package controller

import (
	"go.opentelemetry.io/otel/trace"

	"GoEdu/internal/api/http/server"
	"GoEdu/internal/registration/api/inprocess"
)

type Controller struct {
	*registrationHTTPController
	*userAccessHTTPController
}

var _ server.ServerInterface = (*Controller)(nil)

func New(
	tracer trace.Tracer,
	rAPI inprocess.RegistrationModuleFacade,
) *Controller {
	return &Controller{
		registrationHTTPController: newRegistrationHTTPController(tracer, rAPI),
		userAccessHTTPController:   newUserAccessHTTPController(tracer),
	}
}
