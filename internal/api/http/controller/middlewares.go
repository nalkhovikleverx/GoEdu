package controller

import (
	"log/slog"
	"net/http"

	"GoEdu/internal/api/http/server"
	"GoEdu/internal/pkg/logger"
)

func NewLoggerHTTPMiddleware(parent *slog.Logger) server.MiddlewareFunc {
	return server.MiddlewareFunc(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = logger.WithContext(ctx, slog.New(parent.Handler()))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})
}
