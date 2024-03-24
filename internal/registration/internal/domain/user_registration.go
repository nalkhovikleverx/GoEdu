package domain

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

var (
	UserRegistrationCannotBeConfirmedMoreThanOnceError = fmt.Errorf("UserRegistration can't be confirmed more than once")
)

func RegisterNewUser(firstName, lastName, email, password string) *UserRegistration {
	return &UserRegistration{
		ID:               UserRegistrationId{Value: uuid.New()},
		status:           WaitForConfirmation,
		email:            *CreateUserEmail(email),
		Name:             firstName + " " + lastName,
		FirstName:        firstName,
		LastName:         lastName,
		Password:         password,
		RegistrationDate: time.Now(),
		ConfirmationDate: time.Time{},
	}
}

type UserRegistration struct {
	ID               UserRegistrationId
	status           UserRegistrationStatus
	email            UserRegistrationEmail
	Name             string
	FirstName        string
	LastName         string
	Password         string
	RegistrationDate time.Time
	ConfirmationDate time.Time
}

func (u *UserRegistration) Confirm() {
	u.status = Confirmed
	u.ConfirmationDate = time.Now()
}

type UserRegistrationEvent struct {
	ID               UserRegistrationId
	email            UserRegistrationEmail
	Name             string
	FirstName        string
	LastName         string
	RegistrationDate time.Time
}

type UserRegistrationId struct {
	Value uuid.UUID
}

func CreateUserEmail(value string) *UserRegistrationEmail {
	return &UserRegistrationEmail{value: value}
}

type UserRegistrationEmail struct {
	value string
}

type UserRegistrationStatus string

const (
	WaitForConfirmation UserRegistrationStatus = "wait"
	Confirmed           UserRegistrationStatus = "confirmed"
)
