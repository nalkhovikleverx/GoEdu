package domain_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"GoEdu/internal/registration/internal/domain"
)

func TestPositiveUserPassword(t *testing.T) {
	tests := map[string]struct {
		password string
	}{
		"case 1": {
			password: "aaaa",
		},
		"case 2": {
			password: "1",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			password, err := domain.CreateUserPassword(tc.password)
			require.Nil(t, err)
			require.Equal(t, tc.password, password.String())
		})
	}
}

func TestNegativeUserPassword(t *testing.T) {
	tests := map[string]struct {
		password string
	}{
		"case 1": {
			password: "",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := domain.CreateUserPassword(tc.password)
			require.NotNil(t, err)
		})
	}
}
