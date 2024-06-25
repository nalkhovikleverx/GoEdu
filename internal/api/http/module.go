package http

import (
	"context"

	"GoEdu/internal/api/http/controller"
	"GoEdu/internal/api/http/server"
	"GoEdu/internal/pkg/module"
)

type Module struct{}

func (m *Module) Init(ctx context.Context, dep module.Dependencies) error {
	return Root(ctx, dep)
}

func Root(ctx context.Context, dep module.Dependencies) error {
	mux := dep.HTTP()
	tp := dep.TraceProvider()
	log := dep.Logger()

	impl := controller.New(tp.Tracer("controller"), dep.RegistrationAPI())
	_ = server.HandlerWithOptions(impl, server.StdHTTPServerOptions{
		BaseRouter: mux,
		Middlewares: []server.MiddlewareFunc{
			controller.NewLoggerHTTPMiddleware(log.With("module", "http-api")),
		},
	})
	return nil
}
