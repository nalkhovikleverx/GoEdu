package application

import (
	"GoEdu/internal/registration/internal/domain"
	"context"
	"fmt"
)

type CommandHandler interface {
	Handle(context context.Context, command Command) (CommandResult, error)
}

type RegisterNewUserCommandHandler struct {
	hasher     PasswordHasher
	verifier   UniqueEmailVerifier
	repository UserRegistrationRepository
}

func (r *RegisterNewUserCommandHandler) Handle(context context.Context, command Command) (CommandResult, error) {
	regNewUserCommand := command.(RegisterNewUserCommand)

	user, err := domain.RegisterNewUser(
		regNewUserCommand.FirstName,
		regNewUserCommand.LastName,
		regNewUserCommand.Email,
		r.hasher.Hash(regNewUserCommand.Password),
	)

	if err != nil {
		return nil, err
	}

	if r.verifier.isUnique(user.Email) != true {
		return nil, UserEmailMustBeUniqueError
	}

	err = r.repository.Add(context, *user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

type ConfirmUserRegistrationCommandHandler struct {
	repository UserRegistrationRepository
}

func (c *ConfirmUserRegistrationCommandHandler) Handle(context context.Context, command Command) (CommandResult, error) {
	confUserRegistrationCommand := command.(ConfirmUserRegistrationCommand)
	user, err := c.repository.Load(context, confUserRegistrationCommand.ID)

	if err != nil {
		return nil, err
	}

	if err = user.Confirm(); err != nil {
		return nil, err
	}

	err = c.repository.Update(context, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

var (
	UserEmailMustBeUniqueError = fmt.Errorf("User email must be unique")
)

type UniqueEmailVerifier interface {
	isUnique(email domain.UserRegistrationEmail) bool
}

type PasswordHasher interface {
	Hash(password string) string
}

type UserRegistrationRepository interface {
	Add(context context.Context, userRegistration domain.UserRegistration) error
	Load(context context.Context, id domain.UserRegistrationId) (domain.UserRegistration, error)
	Update(context context.Context, userRegistration domain.UserRegistration) error
}
