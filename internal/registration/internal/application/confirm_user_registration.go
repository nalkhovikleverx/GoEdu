package application

import (
	"context"

	"GoEdu/internal/registration/internal/domain"
)

type ConfirmUserRegistrationCommand struct {
	ID domain.UserRegistrationID
}

func NewConfirmUserRegistrationCommandHandler(repository UserRegistrationRepository) *ConfirmUserRegistrationCommandHandler {
	return &ConfirmUserRegistrationCommandHandler{repository: repository}
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

	err = c.repository.Update(context, &user)

	if err != nil {
		return nil, err
	}

	return RegisterNewUserCommandResult{}, nil
}
