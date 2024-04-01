package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"GoEdu/internal/registration/internal/application"
)

func TestRegisterNewUser(t *testing.T) {
	tests := map[string]struct {
		input []application.RegisterNewUserCommand
		want  bool
	}{
		"ok": {input: []application.RegisterNewUserCommand{
			application.RegisterNewUserCommand{
				FirstName: "Jon",
				LastName:  "Jackson",
				Email:     "jj@gmail.com",
				Password:  "jj",
			},
		},
			want: true,
		},
		"error": {input: []application.RegisterNewUserCommand{
			application.RegisterNewUserCommand{
				FirstName: "Jon",
				LastName:  "Jackson",
				Email:     "jj@gmail.com",
				Password:  "jj",
			},
			application.RegisterNewUserCommand{
				FirstName: "Jon",
				LastName:  "Jackson",
				Email:     "jj@gmail.com",
				Password:  "jj",
			},
		},
			want: false,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			repo := NewMockRepository()
			verifier := UniqueMailVerifier{Repo: *repo}
			rnuh := application.NewRegisterNewUserCommandHandler(&mockHasher{}, repo, &verifier)
			for i, user := range tc.input {
				_, err := rnuh.Handle(context.Background(), user)
				if tc.want || i == 0 {
					assert.ErrorIs(t, err, nil, "Error must be nil in happy case")
				} else {
					assert.EqualError(t, err, "user email must be unique", "We must found error here")
				}
			}

			stored := len(repo.Storage)
			assert.Equal(t, len(tc.input), stored, "Number of stored users not equal to expected number")
		})
	}

}

func TestConfirmUserRegistration(t *testing.T) {
	repo := NewMockRepository()
	verifier := UniqueMailVerifier{Repo: *repo}
	rnuh := application.NewRegisterNewUserCommandHandler(&mockHasher{}, repo, &verifier)
	rnuc := application.RegisterNewUserCommand{
		FirstName: "Jon",
		LastName:  "Jackson",
		Email:     "jj@gmail.com",
		Password:  "jj",
	}

	_, err := rnuh.Handle(context.Background(), rnuc)
	if err != nil {
		assert.True(t, false, "Find some errors:", err)
	}

	stored := len(repo.Storage)
	assert.Equal(t, 1, stored, "Number of stored users not equal to expected number")

	cur := application.NewConfirmUserRegistrationCommandHandler(repo)

	for _, u := range repo.Storage {
		curc := application.ConfirmUserRegistrationCommand{ID: u.ID}

		_, err = cur.Handle(context.Background(), curc)
		if err != nil {
			assert.True(t, false, "Find some errors:", err)
		} else {
			assert.True(t, true)
		}

	}
}
