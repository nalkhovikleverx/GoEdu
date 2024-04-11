package domain

import (
	"errors"
	"strings"
)

var FirstNameCannotBeEmptyError = errors.New("first name is empty")
var LastNameCannotBeEmptyError = errors.New("last name is empty")

type UserName struct {
	firstName string
	lastName  string
}

func CreateUserName(firstName, lastName string) (UserName, error) {
	if len(strings.TrimSpace(firstName)) == 0 {
		return UserName{}, FirstNameCannotBeEmptyError
	}
	if len(strings.TrimSpace(lastName)) == 0 {
		return UserName{}, LastNameCannotBeEmptyError
	}

	return UserName{firstName: firstName, lastName: lastName}, nil
}

func (u UserName) GetFirstName() string {
	return u.firstName
}

func (u UserName) GetLastName() string {
	return u.lastName
}

func (u UserName) GetFullName() string {
	return u.firstName + " " + u.lastName
}
