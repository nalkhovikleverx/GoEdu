package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserRegistration(t *testing.T) {
	tests := map[string]struct {
		userName UserName
		password HashedUserPassword
		email    UserRegistrationEmail
	}{
		"case 1": {
			userName: UserName{
				firstName: "A",
				lastName:  "A",
			},
			password: HashedUserPassword{value: "aaaa"},
			email:    UserRegistrationEmail{value: "aaa@gmail.com"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ur, err := RegisterNewUser(tc.userName, tc.password, tc.email)
			require.Nil(t, err)
			require.Equal(t, tc.password, ur.Password)
			require.Equal(t, tc.userName, ur.UserName)
			require.Equal(t, tc.email, ur.Email)
		})
	}

}
