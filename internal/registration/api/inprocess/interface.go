package inprocess

import (
	"context"

	"github.com/google/uuid"
)

type RegisterNewUserCommand struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type RegisterNewUserCommandResult struct {
}

type ConfirmRegistrationCommand struct {
	RegistrationID uuid.UUID
}

type ConfirmRegistrationCommandResult struct {
}

type RegistrationModuleFacade interface {
	RegisterNewUser(context.Context, RegisterNewUserCommand) (RegisterNewUserCommandResult, error)
	ConfirmRegistration(context.Context, ConfirmRegistrationCommand) (ConfirmRegistrationCommandResult, error)
}
