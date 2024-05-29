package logger

import (
	"context"
	"log/slog"
)

type logCtxKey struct{}

func WithContext(ctx context.Context, l *slog.Logger) context.Context {
	return context.WithValue(ctx, logCtxKey{}, l)
}

func FromContext(ctx context.Context) *slog.Logger {
	l, _ := ctx.Value(logCtxKey{}).(*slog.Logger)
	return l
}
