package domain

import "net/mail"

func CreateUserEmail(value string) (UserRegistrationEmail, error) {
	_, err := mail.ParseAddress(value)
	if err != nil {
		return "", err
	}
	return UserRegistrationEmail(value), nil
}

type UserRegistrationEmail string
