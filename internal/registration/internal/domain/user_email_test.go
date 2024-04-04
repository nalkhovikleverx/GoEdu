package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPositiveUserEmail(t *testing.T) {
	tests := map[string]struct {
		email string
		want  UserRegistrationEmail
	}{
		"case 1": {
			email: "sdsdas@ss.com",
			want:  UserRegistrationEmail{value: "sdsdas@ss.com"},
		},
		"case 2": {
			email: "s@ss.s",
			want:  UserRegistrationEmail{value: "s@ss.s"},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			email, err := CreateUserEmail(tc.email)
			require.Nil(t, err, "email created with errors")
			require.Equal(t, tc.want, *email, "actual email must be equal to expected")
		})
	}
}

func TestNegativeUserEmail(t *testing.T) {
	tests := map[string]struct {
		email string
		want  UserRegistrationEmail
	}{
		"case 1": {
			email: "s@",
			want:  UserRegistrationEmail{},
		},
		"case 2": {
			email: "s@ss.",
			want:  UserRegistrationEmail{},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			email, err := CreateUserEmail(tc.email)
			require.Equal(t, tc.want, *email, "email created without errors")
			require.NotNil(t, err, "email created without errors")
		})
	}
}
