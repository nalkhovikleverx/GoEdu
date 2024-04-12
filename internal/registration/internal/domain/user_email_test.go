package domain_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"GoEdu/internal/registration/internal/domain"
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
			email, err := domain.CreateUserEmail(tc.email)
			require.Nil(t, err)
			require.NotNil(t, email)
			require.Equal(t, tc.email, email.String())
		})
	}
}
