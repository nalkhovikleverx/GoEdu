package domain_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"GoEdu/internal/registration/internal/domain"
)

func TestPositiveCreateUserEmail(t *testing.T) {
	tests := map[string]struct {
		email string
	}{
		"case 1": {
			email: "email@gmail.com",
		},
		"case 2": {
			email: "ss@ss.com",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			email, err := domain.CreateUserEmail(tc.email)
			require.Nil(t, err)
			require.Equal(t, tc.email, email.String())
		})
	}
}

func TestNegativeCreateUserEmail(t *testing.T) {
	tests := map[string]struct {
		email string
	}{
		"case 1": {
			email: "",
		},
		"case 2": {
			email: "@.com",
		},
		"case 3": {
			email: "aa.com",
		},
		"case 4": {
			email: "a @ a.com",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := domain.CreateUserEmail(tc.email)
			require.NotNil(t, err)
		})
	}
}
