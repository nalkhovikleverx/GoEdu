package domain

import "net/mail"

type UserRegistrationEmail struct {
	value string
}

func NewUserEmail(value string) (UserRegistrationEmail, error) {
	_, err := mail.ParseAddress(value)
	if err != nil {
		return UserRegistrationEmail{}, err
	}
	return UserRegistrationEmail{value: value}, nil
}

func MustNewUserEmail(value string) UserRegistrationEmail {
	email, err := NewUserEmail(value)
	if err != nil {
		panic(err)
	}
	return email
}

func (u UserRegistrationEmail) String() string {
	return u.value
}
