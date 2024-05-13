package controller

import (
	"net/http"

	"go.opentelemetry.io/otel/trace"

	"GoEdu/internal/pkg/logger"
)

type userAccessHTTPController struct {
	tracer trace.Tracer
}

func newUserAccessHTTPController(t trace.Tracer) *userAccessHTTPController {
	return &userAccessHTTPController{tracer: t}
}

func (c *userAccessHTTPController) UserAccessLogin(w http.ResponseWriter, r *http.Request) {
	logger.FromContext(r.Context()).Info("user access login")
	w.WriteHeader(http.StatusNotImplemented)
}
