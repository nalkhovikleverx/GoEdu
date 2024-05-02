package domain

import "time"

type UserRegistrationEvent struct {
	id               UserRegistrationID
	email            UserRegistrationEmail
	name             string
	firstName        string
	lastName         string
	registrationDate time.Time
}
