package application

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"GoEdu/internal/registration/internal/domain"
)

func TestConfirmUserRegistration(t *testing.T) {
	command := ConfirmUserRegistrationCommand{
		ID: domain.NewUserRegistrationID(),
	}

	tests := map[string]struct {
		Repository *UserRepositorySpy
		command    ConfirmUserRegistrationCommand
		want       bool
	}{
		"happy path": {
			Repository: NewUserRepositorySpy(nil, &domain.UserRegistration{}, nil),
			command:    command,
			want:       true,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			handler := NewConfirmUserRegistrationCommandHandler(tc.Repository)
			_, err := handler.Handle(context.Background(), tc.command)
			require.Nil(t, err, "NewRegisterNewUserCommandHandler ended with errors")
			require.NotNil(t, tc.Repository.Loaded, "user not loaded")
			require.NotNil(t, tc.Repository.Updated, "user not updated")
		})
	}
}
