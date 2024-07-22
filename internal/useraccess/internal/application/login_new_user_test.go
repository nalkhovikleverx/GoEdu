package application_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"GoEdu/internal/useraccess/internal/application"
	"GoEdu/internal/useraccess/internal/domain"
)

var _ application.UserRepository = (*LoginUserRepoMock)(nil)

type LoginUserRepoMock struct {
	added  bool
	loaded bool
}

func (r *LoginUserRepoMock) Add(_ context.Context, _ *domain.User) error {
	r.added = true
	return nil
}
func (r *LoginUserRepoMock) LoadUserByEmail(
	_ context.Context,
	_ domain.UserEmail) (*domain.User, error) {
	r.loaded = true
	return &domain.User{}, nil
}
func (r LoginUserRepoMock) Update(_ context.Context, _ *domain.User) error {
	return nil
}

func TestPositiveLoginUserAuthentication(t *testing.T) {
	tests := map[string]struct {
		repository *LoginUserRepoMock
		command    application.LoginCommand
	}{
		"successful register user registration": {
			repository: &LoginUserRepoMock{},
			command: application.LoginCommand{
				Email:    domain.MustNewUserEmail("email@mail.ru"),
				Password: domain.MustNewUserPassword("password"),
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			handler := application.NewLoginCommandHandler(tc.repository)
			require.NotNil(t, handler, "handler is nil")
			_, err := handler.Handle(context.Background(), tc.command)
			require.Nil(t, err, "error not nil")
			require.Equal(t, true, tc.repository.loaded, "user registration not loaded")
		})
	}
}
