package controller

import (
	"GoEdu/internal/api/http/server"
	"GoEdu/internal/registration/api/inprocess"
)

func FromRequestToRegisterNewUserCommand(registration server.NewRegistration) inprocess.RegisterNewUserCommand {
	return inprocess.RegisterNewUserCommand{
		FirstName: registration.FirstName,
		LastName:  registration.LastName,
		Email:     string(registration.Email),
	}
}
