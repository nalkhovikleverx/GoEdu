package registration

import (
	"context"

	"GoEdu/internal/pkg/module"
)

type Module struct {
}

func (m *Module) Init(ctx context.Context, dep module.Dependencies) error {
	return Root(ctx, dep)
}

func Root(ctx context.Context, dep module.Dependencies) error {
	return nil
}
