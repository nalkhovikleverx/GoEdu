package domain

import "time"

type UserRegistrationEvent struct {
	ID               UserRegistrationID
	Email            UserRegistrationEmail
	Name             string
	FirstName        string
	LastName         string
	RegistrationDate time.Time
}

// temporarily for ci/cd checks

var _ UserRegistrationEvent = UserRegistrationEvent{}
