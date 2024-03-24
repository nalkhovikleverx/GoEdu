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
		"successful parse urid from decimal": {
			id: "12345678123456781234567812345678",
		},
		"successful parse urid from hex": {
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

func TestNegativeMustParseUserRegistrationID(t *testing.T) {
	tests := map[string]struct {
		id string
	}{
		"unsuccessful parse urid from hex": {
			id: "12af567812fa567812af567812rc5678",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			require.Panics(t, func() {
				domain.MustParseUserRegistrationID(tc.id)
			})
		})
	}
}

func TestPositiveParseUserRegistrationID(t *testing.T) {
	tests := map[string]struct {
		id string
	}{
		"successful parse urid from decimal": {
			id: "12345678123456781234567812345678",
		},
		"successful parse urid from hex": {
			id: "12af567812fa567812af567812bc5678",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			urid, err := domain.ParseUserRegistrationID(tc.id)
			require.NotEmpty(t, urid)
			require.Nil(t, err)
			require.Equal(t, tc.id, strings.ReplaceAll(urid.String(), "-", ""))
		})
	}
}

func TestNegativeParseUserRegistrationID(t *testing.T) {
	tests := map[string]struct {
		id string
	}{
		"unsuccessful empty string": {
			id: "",
		},
		"unsuccessful wrong length": {
			id: "11111",
		},
		"unsuccessful wrong characters": {
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
