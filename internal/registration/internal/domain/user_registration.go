package domain

import (
	"errors"
	"time"
)

var (
	UserRegistrationCannotBeConfirmedMoreThanOnceError = errors.New("user can't be confirmed more than once")
)

func RegisterNewUser(firstName, lastName, password string, email UserRegistrationEmail) (*UserRegistration, error) {

	return &UserRegistration{
		ID:               NewUserRegistrationID(),
		Status:           WaitForConfirmation,
		Email:            email,
		Name:             firstName + " " + lastName,
		FirstName:        firstName,
		LastName:         lastName,
		Password:         password,
		RegistrationDate: time.Now(),
		ConfirmationDate: time.Time{},
	}, nil
}

type UserRegistration struct {
	ID               UserRegistrationID
	Status           UserRegistrationStatus
	Email            UserRegistrationEmail
	Name             string
	FirstName        string
	LastName         string
	Password         string
	RegistrationDate time.Time
	ConfirmationDate time.Time
}

func (u *UserRegistration) Confirm() error {
	if u.Status == Confirmed {
		return UserRegistrationCannotBeConfirmedMoreThanOnceError
	}
	u.Status = Confirmed
	u.ConfirmationDate = time.Now()
	return nil
}
