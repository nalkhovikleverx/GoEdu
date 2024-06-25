package registration

import (
	"context"

	"GoEdu/internal/pkg/module"
	"GoEdu/internal/registration/internal/application"
	"GoEdu/internal/registration/internal/domain"
	"GoEdu/internal/registration/internal/infrastructure/memory"
	"GoEdu/internal/registration/internal/interfaces/inprocess"
)

type Module struct{}

func (m *Module) Init(ctx context.Context, dep module.Dependencies) error {
	return Root(ctx, dep)
}

var _ application.PasswordHasher = (*SimpleHasher)(nil)

type SimpleHasher struct{}

func (s SimpleHasher) Hash(password domain.UserPassword) (domain.HashedUserPassword, error) {
	return domain.NewHashedUserPassword(password), nil
}

var _ application.UniqueEmailVerifier = (*SimpleVerifier)(nil)

type SimpleVerifier struct {
	repo *memory.InProcessRepository
}

func NewVerifier(repo *memory.InProcessRepository) *SimpleVerifier {
	return &SimpleVerifier{
		repo: repo,
	}
}

func (s SimpleVerifier) IsUnique(ctx context.Context, email domain.UserRegistrationEmail) (bool, error) {
	return true, nil
}

func Root(ctx context.Context, dep module.Dependencies) error {
	repo := memory.NewRepository()
	hasher := SimpleHasher{}
	verificator := NewVerifier(repo)
	confirmationHandler := application.NewConfirmUserRegistrationCommandHandler(repo)
	registrationHander := application.NewRegisterNewUserCommandHandler(hasher, repo, verificator)
	moduleImpl := inprocess.NewRegistrationModuleFacade(registrationHander, confirmationHandler)
	dep.SetRegistrationAPI(moduleImpl)
	return nil
}
