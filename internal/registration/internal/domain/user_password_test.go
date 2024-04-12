package domain_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"GoEdu/internal/registration/internal/domain"
)

func TestUserPassword(t *testing.T) {
	tests := map[string]struct {
		password string
	}{
		"case 1": {
			password: "aaaa",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			password, err := domain.CreateUserPassword(tc.password)
			require.Nil(t, err)
			require.NotNil(t, password)
			require.Equal(t, tc.password, password.String())
		})
	}

}
