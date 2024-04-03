package application

import (
	"context"
	"errors"

	"GoEdu/internal/registration/internal/domain"
)

type Command any

type CommandResult any

type CommandHandler interface {
	Handle(context.Context, Command) (CommandResult, error)
}

var (
	UserEmailMustBeUniqueError = errors.New("user email must be unique")
)

type UniqueEmailVerifier interface {
	IsUnique(context.Context, *domain.UserRegistrationEmail) error
}

type PasswordHasher interface {
	Hash(password *domain.UserPassword) (*domain.UserPassword, error)
}

type UserRegistrationRepository interface {
	Add(context.Context, *domain.UserRegistration) error
	Load(context.Context, domain.UserRegistrationID) (*domain.UserRegistration, error)
	Update(context.Context, *domain.UserRegistration) error
}
