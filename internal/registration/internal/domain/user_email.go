package domain

import "net/mail"

type UserRegistrationEmail struct {
	value string
}

func CreateUserEmail(value string) (*UserRegistrationEmail, error) {
	_, err := mail.ParseAddress(value)
	if err != nil {
		return &UserRegistrationEmail{}, err
	}
	return &UserRegistrationEmail{value: value}, nil
}

func (u UserRegistrationEmail) String() string {
	return u.value
}
