package inprocess

import (
	"GoEdu/internal/registration/api/inprocess"
	"GoEdu/internal/registration/internal/application"
	"GoEdu/internal/registration/internal/domain"
)

func FromRequestToRegisterNewUserCommand(
	command inprocess.RegisterNewUserCommand,
) (application.RegisterNewUserCommand, error) {
	email, err := domain.NewUserEmail(command.Email)
	if err != nil {
		return application.RegisterNewUserCommand{}, err
	}

	password, err := domain.NewUserPassword(command.Password)
	if err != nil {
		return application.RegisterNewUserCommand{}, err
	}

	return application.RegisterNewUserCommand{
		FirstName: command.FirstName,
		LastName:  command.LastName,
		Email:     email,
		Password:  password,
	}, nil
}

func FromApplicationToRegisterNewUserCommandResult(
	res application.CommandResult,
) inprocess.RegisterNewUserCommandResult {
	return inprocess.RegisterNewUserCommandResult{}
}

func FromRequestToConfirmUserRegistrationCommand(
	command inprocess.ConfirmRegistrationCommand,
) (application.ConfirmUserRegistrationCommand, error) {
	return application.ConfirmUserRegistrationCommand{
		ID: domain.UserRegistrationID(command.RegistrationID),
	}, nil
}

func FromApplicationToConfirmRegistrationCommandResult(
	res application.CommandResult,
) inprocess.ConfirmRegistrationCommandResult {
	return inprocess.ConfirmRegistrationCommandResult{}
}
