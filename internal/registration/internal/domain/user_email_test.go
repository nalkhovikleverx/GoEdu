package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmail(t *testing.T) {
	tests := map[string]struct {
		email string
	}{
		"case 1": {
			email: "email@gmail.com",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			email, err := CreateUserEmail(tc.email)
			require.Nil(t, err)
			require.NotNil(t, email)
			require.Equal(t, tc.email, email.value)
		})
	}
}
