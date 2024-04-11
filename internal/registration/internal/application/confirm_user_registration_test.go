package application

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"GoEdu/internal/registration/internal/domain"
)

var _ UserRegistrationRepository = (*ConfirmUserRegistrationRepoMock)(nil)

type ConfirmUserRegistrationRepoMock struct {
	loaded  bool
	updated bool
}

func (c ConfirmUserRegistrationRepoMock) Add(_ context.Context, _ *domain.UserRegistration) error {
	return nil
}
func (c *ConfirmUserRegistrationRepoMock) Load(_ context.Context, _ domain.UserRegistrationID) (*domain.UserRegistration, error) {
	c.loaded = true
	return &domain.UserRegistration{}, nil
}
func (c *ConfirmUserRegistrationRepoMock) Update(_ context.Context, _ *domain.UserRegistration) error {
	c.updated = true
	return nil
}

func TestConfirmUserRegistration(t *testing.T) {
	command := ConfirmUserRegistrationCommand{
		ID: domain.NewUserRegistrationID(),
	}

	tests := map[string]struct {
		repository *ConfirmUserRegistrationRepoMock
		command    ConfirmUserRegistrationCommand
	}{
		"happy path": {
			repository: &ConfirmUserRegistrationRepoMock{},
			command:    command,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			handler := NewConfirmUserRegistrationCommandHandler(tc.repository)
			_, err := handler.Handle(context.Background(), tc.command)
			require.Nil(t, err, "NewRegisterNewUserCommandHandler ended with errors")
			require.Equal(t, true, tc.repository.loaded, "user not loaded")
			require.Equal(t, true, tc.repository.updated, "user not updated")
		})
	}
}
