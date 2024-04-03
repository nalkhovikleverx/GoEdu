package application

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"GoEdu/internal/registration/internal/domain"
)

var _ UserRegistrationRepository = (*UserRepositorySpy)(nil)

type UserRepositorySpy struct {
	Added   *domain.UserRegistration
	Loaded  *domain.UserRegistration
	Updated *domain.UserRegistration
}

func (u *UserRepositorySpy) Add(_ context.Context, user *domain.UserRegistration) error {
	u.Added = user
	return nil
}

func (u *UserRepositorySpy) Load(_ context.Context, urid domain.UserRegistrationID) (*domain.UserRegistration, error) {
	return u.Loaded, nil
}

func (u *UserRepositorySpy) Update(_ context.Context, user *domain.UserRegistration) error {
	u.Updated = user
	return nil
}

var _ UniqueEmailVerifier = (*UniqueEmailVerifierSpy)(nil)

type UniqueEmailVerifierSpy struct {
	checked *domain.UserRegistrationEmail
}

func (u *UniqueEmailVerifierSpy) IsUnique(_ context.Context, email *domain.UserRegistrationEmail) error {
	u.checked = email
	return nil
}

var _ PasswordHasher = (*PasswordHasherSpy)(nil)

type PasswordHasherSpy struct {
	hashed *domain.UserPassword
}

func (p *PasswordHasherSpy) Hash(password *domain.UserPassword) (*domain.UserPassword, error) {
	p.hashed = password
	return password, nil
}

func TestRegisterNewUser(t *testing.T) {
	fName, _ := domain.CreateUserFirstName("jon")
	lName, _ := domain.CreateUserLastName("joster")
	email, _ := domain.CreateUserEmail("jojo@gmail.com")
	password, _ := domain.CreateUserPassword("jojo")
	command := RegisterNewUserCommand{
		FirstName: *fName,
		LastName:  *lName,
		Email:     *email,
		Password:  *password,
	}

	tests := map[string]struct {
		Repository *UserRepositorySpy
		Hasher     *PasswordHasherSpy
		verifier   *UniqueEmailVerifierSpy
		command    RegisterNewUserCommand
		want       bool
	}{
		"happy path": {
			Repository: &UserRepositorySpy{
				Added:   nil,
				Loaded:  nil,
				Updated: nil,
			},
			Hasher:   &PasswordHasherSpy{nil},
			verifier: &UniqueEmailVerifierSpy{nil},
			command:  command,
			want:     true,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			handler := NewRegisterNewUserCommandHandler(tc.Hasher, tc.Repository, tc.verifier)
			_, err := handler.Handle(context.Background(), tc.command)
			require.Nil(t, err, "NewRegisterNewUserCommandHandler ended with errors")
			require.NotNil(t, tc.Repository.Added, "user not added to repository")
			require.NotNil(t, tc.Hasher.hashed, "password not hashed")
			require.NotNil(t, tc.verifier.checked, "email not checked")
		})
	}

}
