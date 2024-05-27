package application_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"GoEdu/internal/registration/internal/application"
	"GoEdu/internal/registration/internal/domain"
)

var _ application.UserRegistrationRepository = (*CreateNewUserRepoMock)(nil)

type CreateNewUserRepoMock struct {
	added bool
}

func (r *CreateNewUserRepoMock) Add(_ context.Context, _ *domain.UserRegistration) error {
	r.added = true
	return nil
}
func (r CreateNewUserRepoMock) Load(
	_ context.Context,
	_ domain.UserRegistrationID) (*domain.UserRegistration, error) {
	return &domain.UserRegistration{}, nil
}
func (r CreateNewUserRepoMock) Update(_ context.Context, _ *domain.UserRegistration) error {
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
	return domain.NewHashedUserPassword(password), nil
}

func TestPositiveCreateNewUserRegistration(t *testing.T) {
	tests := map[string]struct {
		verifier   *UniqueEmailVerifierSpy
		hasher     *PasswordHasherSpy
		repository *CreateNewUserRepoMock
		command    application.CreateUserCommand
	}{
		"successful Create user registration": {
			verifier:   &UniqueEmailVerifierSpy{},
			hasher:     &PasswordHasherSpy{},
			repository: &CreateNewUserRepoMock{},
			command: application.CreateUserCommand{
				FirstName: "a",
				LastName:  "a",
				Email:     domain.MustNewUserEmail("email@mail.ru"),
				Password:  domain.MustNewUserPassword("password"),
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			handler := application.NewCreateUserCommandHandler(tc.hasher, tc.repository)
			require.NotNil(t, handler, "handler is nil")
			_, err := handler.Handle(context.Background(), tc.command)
			require.Nil(t, err, "error not nil")
			require.Equal(t, true, tc.hasher.hashed, "password not hashed")
			require.Equal(t, true, tc.repository.added, "user registration not added")
		})
	}
}
