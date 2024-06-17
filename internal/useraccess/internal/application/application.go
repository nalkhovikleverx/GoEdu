package application

import (
	"context"
	"errors"

	"GoEdu/internal/useraccess/internal/domain"
)

type Command any

type CommandResult any

type CommandHandler interface {
	Handle(context.Context, Command) (CommandResult, error)
}

var (
	ErrUserEmailMustBeUnique = errors.New("user email must be unique")
)

type UniqueEmailVerifier interface {
	IsUnique(context.Context, domain.UserRegistrationEmail) error
}

type UniquePasswordVerifier interface {
	IsUnique(context.Context, domain.HashedUserPassword) error
}

type UserCreationRepository interface {
	Add(context.Context, *domain.UserAuthentication) error
	Load(context.Context, domain.UserAuthentication) (*domain.UserAuthentication, error)
	Update(context.Context, *domain.UserAuthentication) error
}
