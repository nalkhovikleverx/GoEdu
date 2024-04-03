package domain

import (
	"errors"
	"strings"
)

var FirstNameCannotBeEmptyError = errors.New("first name is empty")

type UserFirstName struct {
	value string
}

func CreateUserFirstName(value string) (*UserFirstName, error) {
	if len(strings.TrimSpace(value)) != 0 {
		return &UserFirstName{}, FirstNameCannotBeEmptyError
	}
	return &UserFirstName{value: value}, nil
}

func (u UserFirstName) GetValue() string {
	return u.value
}
