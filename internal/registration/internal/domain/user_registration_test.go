package domain_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"GoEdu/internal/registration/internal/domain"
)

func TestUserRegistration(t *testing.T) {
	tests := map[string]struct {
		userName domain.UserName
		password domain.HashedUserPassword
		email    domain.UserRegistrationEmail
	}{
		"successful creation userRegistration": {
			userName: domain.MustNewUserName("A", "A"),
			password: domain.NewHashedUserPassword(domain.MustNewUserPassword("aaaa")),
			email:    domain.MustNewUserEmail("aaa@gmail.com"),
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
