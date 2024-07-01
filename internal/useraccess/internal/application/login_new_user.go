package application

import (
	"context"

	"GoEdu/internal/useraccess/internal/domain"
)

type LoginCommand struct {
	Email    string
	Password string
}

type LoginCommandResult struct {
	UserClaims
}

type UserClaims struct {
	UserID int
	Email  string
}

func NewLoginCommandHandler(repository UserCreationRepository) *LoginCommandHandler {
	return &LoginCommandHandler{repository: repository}
}

type LoginCommandHandler struct {
	repository UserCreationRepository
}

func (r *LoginCommandHandler) Handle(ctx context.Context, command Command) (CommandResult, error) {
	loginCommand := command.(LoginCommand)

	user, err := r.repository.Load(ctx, loginCommand.Email)
	if err != nil {
		return LoginCommandResult{}, err
	}

	if !user.VerifyPassword(loginCommand.Password) {
		return LoginCommandResult{}, errors.New("invalid email or password")
	}

	userClaims := UserClaims{
		UserID: user.ID,
		Email:  user.Email,
	}

	return LoginCommandResult{userClaims}, nil
}
