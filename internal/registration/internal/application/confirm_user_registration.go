package application

import (
	"context"

	"GoEdu/internal/registration/internal/domain"
)

type ConfirmUserRegistrationCommand struct {
	ID domain.UserRegistrationID
}

type ConfirmUserRegistrationCommandHandler struct {
	repository UserRegistrationRepository
}

func NewConfirmUserRegistrationCommandHandler(repository UserRegistrationRepository) *ConfirmUserRegistrationCommandHandler {
	return &ConfirmUserRegistrationCommandHandler{repository: repository}
}

func (c *ConfirmUserRegistrationCommandHandler) Handle(ctx context.Context, command Command) (CommandResult, error) {
	confUserRegistrationCommand := command.(ConfirmUserRegistrationCommand)

	user, err := c.repository.Load(ctx, confUserRegistrationCommand.ID)
	if err != nil {
		return RegisterNewUserCommandResult{}, err
	}

	if err = user.Confirm(); err != nil {
		return RegisterNewUserCommandResult{}, err
	}

	err = c.repository.Update(ctx, user)
	if err != nil {
		return RegisterNewUserCommandResult{}, err
	}

	return RegisterNewUserCommandResult{}, nil
}
