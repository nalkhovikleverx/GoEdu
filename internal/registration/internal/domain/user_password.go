package domain

import (
	"errors"
	"strings"
)

var ErrPasswordCannotBeEmpty = errors.New("password is empty")

type UserPassword struct {
	value string
}

func CreateUserPassword(value string) (UserPassword, error) {
	if len(strings.TrimSpace(value)) == 0 {
		return UserPassword{}, ErrPasswordCannotBeEmpty
	}
	return UserPassword{value}, nil
}

func (u UserPassword) String() string {
	return u.value
}

type HashedUserPassword struct {
	value string
}

func CreateHashedUserPassword(password UserPassword) HashedUserPassword {
	return HashedUserPassword{value: password.String()}
}

func (h HashedUserPassword) String() string {
	return h.value
}
