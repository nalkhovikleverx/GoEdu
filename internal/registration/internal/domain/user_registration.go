package domain

import (
	"errors"
	"net/mail"
	"time"

	"github.com/google/uuid"
)

var (
	UserRegistrationCannotBeConfirmedMoreThanOnceError = errors.New("user can't be confirmed more than once")
)

func RegisterNewUser(firstName, lastName, email, password string) (*UserRegistration, error) {
	userMail, err := CreateUserEmail(email)
	if err != nil {
		return nil, err
	}

	return &UserRegistration{
		ID:               NewUserRegistrationID(),
		Status:           WaitForConfirmation,
		Email:            userMail,
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

type UserRegistrationEvent struct {
	ID               UserRegistrationID
	email            UserRegistrationEmail
	Name             string
	FirstName        string
	LastName         string
	RegistrationDate time.Time
}

func NewUserRegistrationID() UserRegistrationID {
	return UserRegistrationID(uuid.New())
}

type UserRegistrationID uuid.UUID

func ParseUserRegistrationID(urid string) (UserRegistrationID, error) {
	uid, err := uuid.Parse(urid)
	if err != nil {
		return UserRegistrationID{}, err
	}
	return UserRegistrationID(uid), nil
}

func MustParseUserRegistrationID(urid string) UserRegistrationID {
	return UserRegistrationID(uuid.MustParse(urid))
}

func CreateUserEmail(value string) (UserRegistrationEmail, error) {
	_, err := mail.ParseAddress(value)
	if err != nil {
		return UserRegistrationEmail{}, err
	}
	return UserRegistrationEmail{value: value}, nil
}

type UserRegistrationEmail struct {
	value string
}

type UserRegistrationStatus string

const (
	WaitForConfirmation UserRegistrationStatus = "wait"
	Confirmed           UserRegistrationStatus = "confirmed"
)
