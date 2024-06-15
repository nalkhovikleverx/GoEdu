package api

import (
	memory "GoEdu/internal/registration/internal/infrastructure/inprocess/registration"
)

var DefaultRegistrationModuleFacade = memory.CreateNewInProcessRegistrationModuleFacade()
