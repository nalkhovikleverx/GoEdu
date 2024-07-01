package application

import (
	"context"

	"GoEdu/internal/useraccess/internal/domain"
)

type CreateUserCommand struct {
	FirstName string
	LastName  string
	Email     domain.UserRegistrationEmail
	Password  string
}

type CreateUserCommandResult struct {
}

func NewCreateUserCommandHandler(
	repository UserCreationRepository) *CreateUserCommandHandler {
	return &CreateUserCommandHandler{repository: repository}
}

type CreateUserCommandHandler struct {
	repository UserCreationRepository
}

func (r *CreateUserCommandHandler) Handle(ctx context.Context, command Command) (CommandResult, error) {
	createUserCommand := command.(CreateUserCommand)

	hashedPassword := createUserCommand.Password

	userName, err := domain.NewUserName(createUserCommand.FirstName, createUserCommand.LastName)
	if err != nil {
		return CreateUserCommandResult{}, err
	}

	user, err := domain.NewUserRegistration(
		userName,
		hashedPassword,
		createUserCommand.Email,
	)
	if err != nil {
		return CreateUserCommandResult{}, err
	}

	err = r.repository.Add(ctx, user)
	if err != nil {
		return CreateUserCommandResult{}, err
	}

	return CreateUserCommandResult{}, nil
}
