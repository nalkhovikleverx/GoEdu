package application

import (
	"context"

	"GoEdu/internal/registration/internal/domain"
)

type RegisterNewUserCommand struct {
	FirstName string
	LastName  string
	Email     domain.UserRegistrationEmail
	Password  string
}

type RegisterNewUserCommandResult struct {
}

func NewRegisterNewUserCommandHandler(hasher PasswordHasher, repository UserRegistrationRepository, verifier UniqueEmailVerifier) *RegisterNewUserCommandHandler {
	return &RegisterNewUserCommandHandler{hasher: hasher, repository: repository, verifier: verifier}
}

type RegisterNewUserCommandHandler struct {
	hasher     PasswordHasher
	repository UserRegistrationRepository
	verifier   UniqueEmailVerifier
}

func (r *RegisterNewUserCommandHandler) Handle(context context.Context, command Command) (CommandResult, error) {
	regNewUserCommand := command.(RegisterNewUserCommand)

	err := r.verifier.IsUnique(context, regNewUserCommand.Email)
	if err != nil {
		return nil, err
	}

	h, err := r.hasher.Hash(regNewUserCommand.Password)
	if err != nil {
		return nil, err
	}

	user, err := domain.RegisterNewUser(
		regNewUserCommand.FirstName,
		regNewUserCommand.LastName,
		h,
		regNewUserCommand.Email,
	)
	if err != nil {
		return nil, err
	}

	err = r.repository.Add(context, user)
	if err != nil {
		return nil, err
	}

	return RegisterNewUserCommandResult{}, nil
}
