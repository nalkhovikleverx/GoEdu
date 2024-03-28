package application

import "context"

type GetWaitingForConfirmationUserRegistrationsQuery struct {
}

type UserRegistrationReadModel interface {
	GetWaitingForConfirmationUserRegistrations(context.Context, GetWaitingForConfirmationUserRegistrationsQuery) ([]UserRegistrationView, error)
}
