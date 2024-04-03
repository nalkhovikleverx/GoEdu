package domain

import (
	"errors"
	"strings"
)

var LastNameCannotBeEmptyError = errors.New("last name is empty")

type UserLastName struct {
	value string
}

func CreateUserLastName(value string) (*UserLastName, error) {
	if len(strings.TrimSpace(value)) != 0 {
		return &UserLastName{}, LastNameCannotBeEmptyError
	}
	return &UserLastName{value}, nil
}

func (u UserLastName) GetValue() string {
	return u.value
}
