package application

import (
	"context"
	"time"
)

type UserRegistrationView struct {
	Id               string
	FirstName        string
	LastName         string
	Email            string
	Status           string
	RegistrationDate time.Time
}

type GetWaitingForConfirmationUserRegistrationsQuery struct {
}

type UserRegistrationReadModel interface {
	GetWaitingForConfirmationUserRegistrations(context.Context, GetWaitingForConfirmationUserRegistrationsQuery) ([]UserRegistrationView, error)
}
