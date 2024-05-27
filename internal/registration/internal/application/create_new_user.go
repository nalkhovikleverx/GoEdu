package application

import (
	"context"

	"GoEdu/internal/registration/internal/domain"
)

type CreateUserCommand struct {
	FirstName string
	LastName  string
	Email     domain.UserRegistrationEmail
	Password  domain.UserPassword
}

type CreateUserCommandResult struct {
}

func NewCreateUserCommandHandler(
	hasher PasswordHasher,
	repository UserRegistrationRepository) *CreateUserCommandHandler {
	return &CreateUserCommandHandler{hasher: hasher, repository: repository}
}

type CreateUserCommandHandler struct {
	hasher     PasswordHasher
	repository UserRegistrationRepository
}

func (r *CreateUserCommandHandler) Handle(ctx context.Context, command Command) (CommandResult, error) {
	createUserCommand := command.(CreateUserCommand)

	hashedPassword, err := r.hasher.Hash(createUserCommand.Password)
	if err != nil {
		return CreateUserCommandResult{}, err
	}

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

	if err = user.Confirm(); err != nil {
		return CreateUserCommandResult{}, err
	}

	err = r.repository.Add(ctx, user)
	if err != nil {
		return CreateUserCommandResult{}, err
	}

	return CreateUserCommandResult{}, nil
}
