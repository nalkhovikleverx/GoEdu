package domain_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"GoEdu/internal/registration/internal/domain"
)

func TestPositiveUserName(t *testing.T) {
	tests := map[string]struct {
		firstName string
		lastName  string
	}{
		"case 1": {
			firstName: "A",
			lastName:  "B",
		},
		"case 2": {
			"jotaro",
			"joske",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			uName, err := domain.CreateUserName(tc.firstName, tc.lastName)
			require.Nil(t, err)
			require.Equal(t, tc.firstName, uName.GetFirstName())
			require.Equal(t, tc.lastName, uName.GetLastName())
			require.Equal(t, tc.firstName+" "+tc.lastName, uName.GetFullName())
		})
	}
}

func TestNegativeUserName(t *testing.T) {
	tests := map[string]struct {
		firstName string
		lastName  string
	}{
		"case 1": {
			firstName: "",
			lastName:  "B",
		},
		"case 2": {
			firstName: "A",
			lastName:  "",
		},
		"case 3": {
			firstName: "",
			lastName:  "",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := domain.CreateUserName(tc.firstName, tc.lastName)
			require.NotNil(t, err)
		})
	}
}
