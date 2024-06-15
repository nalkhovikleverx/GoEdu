package memory

import (
	"context"

	"GoEdu/internal/registration/api/registration"
	"GoEdu/internal/registration/internal/application"
	"GoEdu/internal/registration/internal/domain"
	"GoEdu/internal/registration/internal/infrastructure/inprocess/repository/memory"
)

var _ registration.RegistrationModuleFacade = (*RegistrationModuleFacade)(nil)

type RegistrationModuleFacade struct {
	registrationHandler application.CommandHandler
	confirmationHandler application.CommandHandler
}

func NewRegistrationModuleFacade(
	registrationHandler,
	confirmationHandler application.CommandHandler,
) *RegistrationModuleFacade {
	return &RegistrationModuleFacade{
		registrationHandler: registrationHandler,
		confirmationHandler: confirmationHandler,
	}
}

func CreateNewInProcessRegistrationModuleFacade() *RegistrationModuleFacade {
	repo := memory.NewRepository()
	hasher := SimpleHasher{}
	verificator := NewVerifier(repo)
	confirmationHandler := application.NewConfirmUserRegistrationCommandHandler(repo)
	registrationHander := application.NewRegisterNewUserCommandHandler(hasher, repo, verificator)

	return NewRegistrationModuleFacade(registrationHander, confirmationHandler)
}

func (r RegistrationModuleFacade) RegisterNewUser(
	ctx context.Context,
	command registration.RegisterNewUserCommand,
) (registration.RegisterNewUserCommandResult, error) {
	res, err := r.registrationHandler.Handle(ctx, command)
	return res.(registration.RegisterNewUserCommandResult), err
}

func (r RegistrationModuleFacade) ConfirmRegistration(
	ctx context.Context,
	command registration.ConfirmRegistrationCommand,
) (registration.ConfirmRegistrationCommandResult, error) {
	res, err := r.confirmationHandler.Handle(ctx, command)
	return res.(registration.ConfirmRegistrationCommandResult), err
}

var _ application.PasswordHasher = (*SimpleHasher)(nil)

type SimpleHasher struct {
}

func (s SimpleHasher) Hash(password domain.UserPassword) (domain.HashedUserPassword, error) {
	return domain.NewHashedUserPassword(password), nil
}

var _ application.UniqueEmailVerifier = (*SimpleVerifier)(nil)

type SimpleVerifier struct {
	repo application.UserRegistrationRepository
}

func NewVerifier(repo application.UserRegistrationRepository) *SimpleVerifier {
	return &SimpleVerifier{
		repo: repo,
	}
}

func (s SimpleVerifier) IsUnique(ctx context.Context, email domain.UserRegistrationEmail) (bool, error) {
	all := s.repo.GetAll(ctx)
	for _, v := range all {
		if v.GetSnapshot().Email == email {
			return false, application.ErrUserEmailMustBeUnique
		}
	}
	return true, nil
}
