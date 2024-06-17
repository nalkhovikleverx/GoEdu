package domain

import "time"

type UserRegistrationEvent struct {
	ID               string
	Email            string
	Name             string
	FirstName        string
	LastName         string
	RegistrationDate time.Time
}
