package domain

type UserRegistrationStatus string

const (
	WaitForConfirmation UserRegistrationStatus = "wait"
	Confirmed           UserRegistrationStatus = "confirmed"
)
