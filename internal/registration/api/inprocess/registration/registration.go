package inprocess

import (
	"context"

	"GoEdu/internal/registration/api/registration"
	"GoEdu/internal/registration/internal/application"
	"GoEdu/internal/registration/internal/domain"
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
	repo := NewRepository()
	hasher := SimpleHasher{}
	verificator := SimpleVerifier{repo: repo}
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
	repo *InProcessRepository
}

func (s SimpleVerifier) IsUnique(ctx context.Context, email domain.UserRegistrationEmail) error {
	for _, u := range s.repo.store {
		if u.GetSnapshot().Email == email {
			return application.ErrUserEmailMustBeUnique
		}
	}
	return nil
}
