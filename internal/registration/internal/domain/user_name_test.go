package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserName(t *testing.T) {
	tests := map[string]struct {
		firstName string
		lastName  string
	}{
		"case 1": {
			firstName: "A",
			lastName:  "B",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			uName, err := CreateUserName(tc.firstName, tc.lastName)
			require.Nil(t, err)
			require.NotNil(t, uName)
			require.Equal(t, tc.firstName, uName.GetFirstName())
			require.Equal(t, tc.lastName, uName.GetLastName())
			require.Equal(t, tc.firstName+" "+tc.lastName, uName.GetFullName())
		})
	}
}
