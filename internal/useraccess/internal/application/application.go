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
	ErrNotFound              = errors.New("not found")
)

type PasswordManager interface {
	IsEqual(domain.UserPassword, domain.HashedUserPassword) bool
}

type UserRepository interface {
	Add(context.Context, *domain.User) error
	LoadUserByEmail(context.Context, domain.UserEmail) (*domain.User, error)
	Update(context.Context, *domain.User) error
}
