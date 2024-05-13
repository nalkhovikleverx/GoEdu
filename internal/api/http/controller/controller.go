package controller

import (
	"go.opentelemetry.io/otel/trace"

	"GoEdu/internal/api/http/server"
)

type Controller struct {
	*registrationHTTPController
	*userAccessHTTPController
}

var _ server.ServerInterface = (*Controller)(nil)

func New(tracer trace.Tracer) *Controller {
	return &Controller{
		registrationHTTPController: newRegistrationHTTPController(tracer),
		userAccessHTTPController:   newUserAccessHTTPController(tracer),
	}
}
