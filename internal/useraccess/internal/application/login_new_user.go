package application

import (
	"context"
	"errors"

	"GoEdu/internal/useraccess/internal/domain"
)

type LoginCommand struct {
	Email    domain.UserEmail
	Password domain.UserPassword
}

type LoginCommandResult struct {
	userClaims UserClaims
}

type UserClaims struct {
	UserID domain.UserID
	Email  domain.UserEmail
}

func NewLoginCommandHandler(repository UserRepository, manager PasswordManager) *LoginCommandHandler {
	return &LoginCommandHandler{
		repository:      repository,
		passwordManager: manager,
	}
}

type LoginCommandHandler struct {
	repository      UserRepository
	passwordManager PasswordManager
}

func (r *LoginCommandHandler) Handle(ctx context.Context, command Command) (CommandResult, error) {
	loginCommand := command.(LoginCommand)

	user, err := r.repository.LoadUserByEmail(ctx, loginCommand.Email)
	if err != nil {
		return LoginCommandResult{}, err
	}

	if !r.passwordManager.IsEqual(loginCommand.Password, user.GetPassword()) {
		return LoginCommandResult{}, errors.New("invalid email or password")
	}

	userClaims := UserClaims{
		UserID: user.GetID(),
		Email:  user.GetEmail(),
	}

	return LoginCommandResult{userClaims}, nil
}
