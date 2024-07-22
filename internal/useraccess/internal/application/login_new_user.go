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

func NewLoginCommandHandler(repository UserRepository) *LoginCommandHandler {
	return &LoginCommandHandler{
		repository: repository,
	}
}

type LoginCommandHandler struct {
	repository UserRepository
}

func (r *LoginCommandHandler) Handle(ctx context.Context, command Command) (CommandResult, error) {
	loginCommand := command.(LoginCommand)

	user, err := r.repository.LoadUserByEmail(ctx, loginCommand.Email)
	if err != nil {
		return LoginCommandResult{}, err
	}

	userS := user.GetUserSnapshot()
	if user.IsPasswordEqual(loginCommand.Password) {
		return LoginCommandResult{}, errors.New("invalid email or password")
	}

	userClaims := UserClaims{
		UserID: userS.ID,
		Email:  userS.Email,
	}

	return LoginCommandResult{userClaims}, nil
}
