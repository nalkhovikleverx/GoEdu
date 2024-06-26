package inprocess

import (
	"context"

	"GoEdu/internal/registration/api/inprocess"
	"GoEdu/internal/registration/internal/application"
)

var _ inprocess.RegistrationModuleFacade = (*inprocessModuleFacade)(nil)

type inprocessModuleFacade struct {
	inprocessHandler    application.CommandHandler
	confirmationHandler application.CommandHandler
}

func NewRegistrationModuleFacade(
	inprocessHandler,
	confirmationHandler application.CommandHandler,
) *inprocessModuleFacade {
	return &inprocessModuleFacade{
		inprocessHandler:    inprocessHandler,
		confirmationHandler: confirmationHandler,
	}
}

func (r inprocessModuleFacade) RegisterNewUser(
	ctx context.Context,
	command inprocess.RegisterNewUserCommand,
) (inprocess.RegisterNewUserCommandResult, error) {
	cmd, err := FromRequestToRegisterNewUserCommand(command)
	if err != nil {
		return inprocess.RegisterNewUserCommandResult{}, err
	}

	res, err := r.inprocessHandler.Handle(ctx, cmd)
	return FromApplicationToRegisterNewUserCommandResult(res), err
}

func (r inprocessModuleFacade) ConfirmRegistration(
	ctx context.Context,
	command inprocess.ConfirmRegistrationCommand,
) (inprocess.ConfirmRegistrationCommandResult, error) {
	cmd, err := FromRequestToConfirmUserRegistrationCommand(command)
	if err != nil {
		return inprocess.ConfirmRegistrationCommandResult{}, err
	}

	res, err := r.confirmationHandler.Handle(ctx, cmd)
	return FromApplicationToConfirmRegistrationCommandResult(res), err
}
