package application

import (
	"GoEdu/internal/registration/internal/domain"
	"fmt"
)

// Event

type Event interface {
}

// Command

type Command interface {
}

type CommandResult struct {
	Result Result
	Events []Event
}

type RegisterNewUserCommand struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type ConfirmUserRegistrationCommand struct {
	ID domain.UserRegistrationId
}

// UserRegistration

type GetWaitingForConfirmationUserRegistrationsQuery struct {
}

type UserRegistrationReadModel interface {
	GetWaitingForConfirmationUserRegistrations(GetWaitingForConfirmationUserRegistrationsQuery) []domain.UserRegistration
}

//Result

type Result interface {
}

type NoResult struct {
}

// Command Handler

type CommandHandler interface {
	Handle(command Command) []CommandResult
}

type RegisterNewUserCommandHandler struct {
}

func (r *RegisterNewUserCommandHandler) Handle(command Command) []CommandResult {
	panic("Not implemented yet")
}

type ConfirmUserRegistrationCommandHandler struct {
}

func (h *ConfirmUserRegistrationCommandHandler) Handle(command Command) []CommandResult {
	panic("Not implemented yet")
}

var (
	UserEmailMustBeUniqueError = fmt.Errorf("User email must be unique ")
)

type UniqueEmailVerifier interface {
	isUnique(email domain.UserRegistrationEmail) bool
}

type PasswordHasher interface {
	Hash(password string) string
}

type UserRegistrationRepository interface {
	Add(userRegistration domain.UserRegistration)
	Load(id domain.UserRegistrationId) domain.UserRegistration
	Update(userRegistration domain.UserRegistration)
}
