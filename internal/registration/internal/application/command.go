package application

import (
	"context"
	"errors"

	"GoEdu/internal/registration/internal/domain"
)

type Command interface {
}

type CommandResult interface {
}

type RegisterNewUserCommand struct {
	FirstName     string
	LastName      string
	VerifiedEmail string
	Password      string
}

type ConfirmUserRegistrationCommand struct {
	ID domain.UserRegistrationID
}

type CommandHandler interface {
	Handle(context.Context, Command) (CommandResult, error)
}

type RegisterNewUserCommandResult struct {
}

type RegisterNewUserCommandHandler struct {
	Hasher     PasswordHasher
	Repository UserRegistrationRepository
}

func (r *RegisterNewUserCommandHandler) Handle(context context.Context, command Command) (CommandResult, error) {
	regNewUserCommand := command.(RegisterNewUserCommand)

	h, err := r.Hasher.Hash(regNewUserCommand.Password)
	if err != nil {
		return nil, err
	}

	user, err := domain.RegisterNewUser(
		regNewUserCommand.FirstName,
		regNewUserCommand.LastName,
		regNewUserCommand.VerifiedEmail,
		h,
	)
	if err != nil {
		return nil, err
	}

	err = r.Repository.Add(context, *user)
	if err != nil {
		return nil, err
	}

	return RegisterNewUserCommandResult{}, nil
}

type ConfirmUserRegistrationCommandHandler struct {
	Repository UserRegistrationRepository
}

func (c *ConfirmUserRegistrationCommandHandler) Handle(context context.Context, command Command) (CommandResult, error) {
	confUserRegistrationCommand := command.(ConfirmUserRegistrationCommand)
	user, err := c.Repository.Load(context, confUserRegistrationCommand.ID)

	if err != nil {
		return nil, err
	}

	if err = user.Confirm(); err != nil {
		return nil, err
	}

	err = c.Repository.Update(context, user)

	if err != nil {
		return nil, err
	}

	return RegisterNewUserCommandResult{}, nil
}

var (
	UserEmailMustBeUniqueError = errors.New("user email must be unique")
)

type UniqueEmailVerifier interface {
	IsUnique(context.Context, domain.UserRegistrationEmail) (bool, error)
}

type PasswordHasher interface {
	Hash(string) (string, error)
}

type UserRegistrationRepository interface {
	Add(context.Context, domain.UserRegistration) error
	Load(context.Context, domain.UserRegistrationID) (domain.UserRegistration, error)
	Update(context.Context, domain.UserRegistration) error
}
