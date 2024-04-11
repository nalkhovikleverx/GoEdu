package domain

import (
	"errors"
	"strings"
)

var PasswordCannotBeEmptyError = errors.New("password is empty")

type UserPassword struct {
	value string
}

func CreateUserPassword(value string) (UserPassword, error) {
	if len(strings.TrimSpace(value)) == 0 {
		return UserPassword{}, PasswordCannotBeEmptyError
	}
	return UserPassword{value}, nil
}

type HashedUserPassword struct {
	value string
}

func CreateHashedUserPassword(password UserPassword) HashedUserPassword {
	return HashedUserPassword{value: password.value}
}
