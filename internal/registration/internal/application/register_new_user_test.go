package application

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"GoEdu/internal/registration/internal/domain"
)

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
			Repository: NewUserRepositorySpy(nil, nil, nil),
			Hasher:     &PasswordHasherSpy{nil},
			verifier:   &UniqueEmailVerifierSpy{nil},
			command:    command,
			want:       true,
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
