package domain_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"GoEdu/internal/registration/internal/domain"
)

func TestPositiveMustParseUserRegistrationID(t *testing.T) {
	tests := map[string]struct {
		id string
	}{
		"case 1": {
			id: "12345678123456781234567812345678",
		},
		"case 2": {
			id: "12af567812fa567812af567812bc5678",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			urid := domain.MustParseUserRegistrationID(tc.id)
			require.NotEmpty(t, urid)
			require.Equal(t, tc.id, strings.ReplaceAll(urid.String(), "-", ""))
		})
	}
}

func TestNegativeParseUserRegistrationID(t *testing.T) {
	tests := map[string]struct {
		id string
	}{
		"case 1": {
			id: "",
		},
		"case 2": {
			id: "11111",
		},
		"case 3": {
			id: "11dffrt",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := domain.ParseUserRegistrationID(tc.id)
			require.NotNil(t, err)
		})
	}
}
