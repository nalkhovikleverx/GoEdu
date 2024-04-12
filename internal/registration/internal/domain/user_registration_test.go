package domain_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"GoEdu/internal/registration/internal/domain"
)

func TestUserRegistration(t *testing.T) {
	user, _ := domain.CreateUserName("A", "A")
	p, _ := domain.CreateUserPassword("aaaa")
	hp := domain.CreateHashedUserPassword(p)
	email, _ := domain.CreateUserEmail("aaa@gmail.com")
	tests := map[string]struct {
		userName domain.UserName
		password domain.HashedUserPassword
		email    domain.UserRegistrationEmail
	}{
		"case 1": {
			userName: user,
			password: hp,
			email:    *email,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ur, err := domain.RegisterNewUser(tc.userName, tc.password, tc.email)
			require.Nil(t, err)
			require.Equal(t, tc.password, ur.Password)
			require.Equal(t, tc.userName, ur.UserName)
			require.Equal(t, tc.email, ur.Email)
		})
	}
}
