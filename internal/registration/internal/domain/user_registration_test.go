package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPositiveUserRegistration(t *testing.T) {
	tests := map[string]struct {
		fName    string
		lName    string
		name     string
		password string
		email    string
		want     UserRegistration
	}{
		"case 1": {
			fName:    "fName",
			lName:    "lName",
			name:     "fName" + " " + "lName",
			password: "1234",
			email:    "email@email.com",
			want: UserRegistration{
				ID:               UserRegistrationID{},
				Status:           "",
				Email:            UserRegistrationEmail{"email@email.com"},
				Name:             "fName" + " " + "lName",
				FirstName:        UserFirstName{"fName"},
				LastName:         UserLastName{"lName"},
				Password:         UserPassword{"1234"},
				RegistrationDate: time.Time{},
				ConfirmationDate: time.Time{},
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			fName, err := CreateUserFirstName(tc.fName)
			require.Nil(t, err, err)
			require.Equal(t, fName.value, tc.fName)
			lName, err := CreateUserLastName(tc.lName)
			require.Nil(t, err, err)
			require.Equal(t, lName.value, tc.lName)
			password, err := CreateUserPassword(tc.password)
			require.Nil(t, err, err)
			require.Equal(t, password.value, tc.password)
			email, err := CreateUserEmail(tc.email)
			require.Nil(t, err, err)
			require.Equal(t, email.value, tc.email)
			ur, err := RegisterNewUser(*fName, *lName, *password, *email)
			require.Nil(t, err, err)
			require.Equal(t, ur.Name, tc.name)

		})
	}
}

func TestNegativeUserRegistration(t *testing.T) {
	tests := map[string]struct {
		fName         string
		lName         string
		password      string
		email         string
		fNameError    error
		lNameError    error
		passwordError error
	}{
		"case 1": {
			fName:         "",
			lName:         "",
			password:      "",
			email:         "",
			fNameError:    FirstNameCannotBeEmptyError,
			lNameError:    LastNameCannotBeEmptyError,
			passwordError: PasswordCannotBeEmptyError,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			fName, err := CreateUserFirstName(tc.fName)
			require.NotNil(t, err, err)
			require.EqualError(t, err, tc.fNameError.Error())
			require.Equal(t, *fName, UserFirstName{})
			lName, err := CreateUserLastName(tc.lName)
			require.NotNil(t, err, err)
			require.EqualError(t, err, tc.lNameError.Error())
			require.Equal(t, *lName, UserLastName{})
			password, err := CreateUserPassword(tc.password)
			require.NotNil(t, err, err)
			require.EqualError(t, err, tc.passwordError.Error())
			require.Equal(t, *password, UserPassword{})
			email, err := CreateUserEmail(tc.email)
			require.NotNil(t, err, err)
			require.Error(t, err)
			require.Equal(t, *email, UserRegistrationEmail{})
		})
	}
}
