package domain

import (
	"errors"
	"strings"
)

var PasswordCannotBeEmptyError = errors.New("password is empty")

type UserPassword struct {
	value string
}

func CreateUserPassword(value string) (*UserPassword, error) {
	if len(strings.TrimSpace(value)) != 0 {
		return &UserPassword{}, PasswordCannotBeEmptyError
	}
	return &UserPassword{value}, nil
}

func (u UserPassword) GetValue() string {
	return u.value
}
