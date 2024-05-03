package domain

import (
	"errors"
	"time"
)

var (
	ErrUserRegistrationCannotBeConfirmedMoreThanOnce = errors.New("user can't be confirmed more than once")
)

type UserRegistration struct {
	id               UserRegistrationID
	status           UserRegistrationStatus
	email            UserRegistrationEmail
	userName         UserName
	password         HashedUserPassword
	registrationDate time.Time
	confirmationDate time.Time
}

func NewUserRegistration(
	userName UserName,
	password HashedUserPassword,
	email UserRegistrationEmail,
) (*UserRegistration, error) {
	return &UserRegistration{
		NewUserRegistrationID(),
		WaitForConfirmation,
		email,
		userName,
		password,
		time.Now(),
		time.Time{},
	}, nil
}

func MustCreateUserRegistrationFromSnapshot(snapshot UserRegistrationSnapshot) *UserRegistration {
	return &UserRegistration{
		snapshot.ID,
		snapshot.Status,
		snapshot.Email,
		snapshot.UserName,
		snapshot.Password,
		snapshot.RegistrationDate,
		snapshot.ConfirmationDate,
	}
}

func (u *UserRegistration) Confirm() error {
	if u.status == Confirmed {
		return ErrUserRegistrationCannotBeConfirmedMoreThanOnce
	}
	u.status = Confirmed
	u.confirmationDate = time.Now()
	return nil
}

func (u *UserRegistration) GetSnapshot() UserRegistrationSnapshot {
	return MustCreateUserRegistrationSnapshot(
		u.id,
		u.status,
		u.email,
		u.userName,
		u.password,
		u.registrationDate,
		u.confirmationDate,
	)
}

type UserRegistrationSnapshot struct {
	ID               UserRegistrationID
	Status           UserRegistrationStatus
	Email            UserRegistrationEmail
	UserName         UserName
	Password         HashedUserPassword
	RegistrationDate time.Time
	ConfirmationDate time.Time
}

func MustCreateUserRegistrationSnapshot(
	id UserRegistrationID,
	status UserRegistrationStatus,
	email UserRegistrationEmail,
	name UserName,
	password HashedUserPassword,
	registrationDate time.Time,
	confirmationDate time.Time,

) UserRegistrationSnapshot {
	return UserRegistrationSnapshot{
		id,
		status,
		email,
		name,
		password,
		registrationDate,
		confirmationDate,
	}
}
