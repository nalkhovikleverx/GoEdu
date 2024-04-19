package reg

import (
	"GoEdu/internal/registration/internal/application"
)

type SomeType struct{}

var _ application.Command = (*SomeType)(nil)
