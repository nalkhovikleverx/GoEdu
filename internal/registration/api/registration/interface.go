package registration

import (
	"context"

	"github.com/google/uuid"
)

type RegisterNewUserCommand struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
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
