package domain

import (
	"errors"
	"time"
)

var (
	UserRegistrationCannotBeConfirmedMoreThanOnceError = errors.New("user can't be confirmed more than once")
)

type UserRegistration struct {
	ID               UserRegistrationID
	Status           UserRegistrationStatus
	Email            UserRegistrationEmail
	Name             string
	FirstName        UserFirstName
	LastName         UserLastName
	Password         UserPassword
	RegistrationDate time.Time
	ConfirmationDate time.Time
}

func RegisterNewUser(
	firstName UserFirstName,
	lastName UserLastName,
	password UserPassword,
	email UserRegistrationEmail) (*UserRegistration, error) {

	return &UserRegistration{
		ID:               NewUserRegistrationID(),
		Status:           WaitForConfirmation,
		Email:            email,
		Name:             firstName.GetValue() + " " + lastName.GetValue(),
		FirstName:        firstName,
		LastName:         lastName,
		Password:         password,
		RegistrationDate: time.Now(),
		ConfirmationDate: time.Time{},
	}, nil
}

func (u *UserRegistration) Confirm() error {
	if u.Status == Confirmed {
		return UserRegistrationCannotBeConfirmedMoreThanOnceError
	}
	u.Status = Confirmed
	u.ConfirmationDate = time.Now()
	return nil
}
