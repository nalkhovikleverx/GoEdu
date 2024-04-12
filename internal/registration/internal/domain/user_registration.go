package domain

import (
	"errors"
	"time"
)

var (
	ErrUserRegistrationCannotBeConfirmedMoreThanOnce = errors.New("user can't be confirmed more than once")
)

type UserRegistration struct {
	ID               UserRegistrationID
	Status           UserRegistrationStatus
	Email            UserRegistrationEmail
	UserName         UserName
	Password         HashedUserPassword
	RegistrationDate time.Time
	ConfirmationDate time.Time
}

func RegisterNewUser(
	userName UserName,
	password HashedUserPassword,
	email UserRegistrationEmail) (*UserRegistration, error) {
	return &UserRegistration{
		ID:               NewUserRegistrationID(),
		Status:           WaitForConfirmation,
		Email:            email,
		UserName:         userName,
		Password:         password,
		RegistrationDate: time.Now(),
		ConfirmationDate: time.Time{},
	}, nil
}

func (u *UserRegistration) Confirm() error {
	if u.Status == Confirmed {
		return ErrUserRegistrationCannotBeConfirmedMoreThanOnce
	}
	u.Status = Confirmed
	u.ConfirmationDate = time.Now()
	return nil
}
