package domain

import "time"

type UserRegistrationEvent struct {
	ID               UserRegistrationID
	email            UserRegistrationEmail
	Name             string
	FirstName        string
	LastName         string
	RegistrationDate time.Time
}
