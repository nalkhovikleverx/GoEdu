package application_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"GoEdu/internal/registration/internal/application"
	"GoEdu/internal/registration/internal/domain"
)

var _ application.UserRegistrationRepository = (*RegisterNewUserRepoMock)(nil)

type RegisterNewUserRepoMock struct {
	added bool
}

func (r *RegisterNewUserRepoMock) Add(_ context.Context, _ *domain.UserRegistration) error {
	r.added = true
	return nil
}
func (r RegisterNewUserRepoMock) Load(_ context.Context, _ domain.UserRegistrationID) (*domain.UserRegistration, error) {
	return nil, nil
}
func (r RegisterNewUserRepoMock) Update(_ context.Context, _ *domain.UserRegistration) error {
	return nil
}

var _ application.UniqueEmailVerifier = (*UniqueEmailVerifierSpy)(nil)

type UniqueEmailVerifierSpy struct {
	checked bool
}

func (u *UniqueEmailVerifierSpy) IsUnique(_ context.Context, _ domain.UserRegistrationEmail) error {
	u.checked = true
	return nil
}

var _ application.PasswordHasher = (*PasswordHasherSpy)(nil)

type PasswordHasherSpy struct {
	hashed bool
}

func (p *PasswordHasherSpy) Hash(password domain.UserPassword) (domain.HashedUserPassword, error) {
	p.hashed = true
	return domain.CreateHashedUserPassword(password), nil
}

func TestRegisterNewUserRegistration(t *testing.T) {
	tests := map[string]struct {
		verifier   *UniqueEmailVerifierSpy
		hasher     *PasswordHasherSpy
		repository *RegisterNewUserRepoMock
		command    application.RegisterNewUserCommand
	}{
		"happy path": {
			verifier:   &UniqueEmailVerifierSpy{},
			hasher:     &PasswordHasherSpy{},
			repository: &RegisterNewUserRepoMock{},
			command: application.RegisterNewUserCommand{
				FirstName: "a",
				LastName:  "a",
				Email:     domain.UserRegistrationEmail{},
				Password:  domain.UserPassword{},
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			handler := application.NewRegisterNewUserCommandHandler(tc.hasher, tc.repository, tc.verifier)
			require.NotNil(t, handler, "handler is nil")
			command, err := handler.Handle(context.Background(), tc.command)
			require.Nil(t, err, "error not nil")
			require.NotNil(t, command, "result is nil")
			require.Equal(t, true, tc.verifier.checked, "verifier not checked")
			require.Equal(t, true, tc.hasher.hashed, "password not hashed")
			require.Equal(t, true, tc.repository.added, "user registration not added")
		})
	}
}
