package test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"GoEdu/internal/registration/internal/application"
	"GoEdu/internal/registration/internal/domain"
)

func NewMockRepository() mockRepository {
	return mockRepository{Storage: make(map[domain.UserRegistrationID]domain.UserRegistration)}
}

type mockRepository struct {
	Storage map[domain.UserRegistrationID]domain.UserRegistration
}

func (m mockRepository) Add(ctx context.Context, registration domain.UserRegistration) error {
	m.Storage[registration.ID] = registration
	return nil
}

func (m mockRepository) Load(ctx context.Context, urid domain.UserRegistrationID) (domain.UserRegistration, error) {
	user, ok := m.Storage[urid]
	if ok != true {
		return domain.UserRegistration{}, errors.New("user not found")
	}
	return user, nil
}

func (m mockRepository) Update(ctx context.Context, user domain.UserRegistration) error {
	_, ok := m.Storage[user.ID]
	if ok != true {
		return errors.New("user not found")
	}
	m.Storage[user.ID] = user
	return nil
}

type mockHasher struct {
}

func (m *mockHasher) Hash(password string) (string, error) {
	return password + "aaa", nil
}

func TestRegisterNewUser(t *testing.T) {
	rnuh := application.RegisterNewUserCommandHandler{Hasher: &mockHasher{}, Repository: NewMockRepository()}
	rnuc := application.RegisterNewUserCommand{
		FirstName:     "Jon",
		LastName:      "Jackson",
		VerifiedEmail: "jj@gmail.com",
		Password:      "jj",
	}

	_, err := rnuh.Handle(context.Background(), rnuc)
	if err != nil {
		t.Errorf("Find some errors %q", err)
	}

	stored := len(rnuh.Repository.(mockRepository).Storage)
	fmt.Println("Stored users :", stored)
	if stored != 1 {
		t.Errorf("Number of stored users %v not equal to expected number %v", stored, 1)
	}
}

func TestConfirmUserRegistration(t *testing.T) {
	repo := NewMockRepository()
	rnuh := application.RegisterNewUserCommandHandler{Hasher: &mockHasher{}, Repository: repo}
	rnuc := application.RegisterNewUserCommand{
		FirstName:     "Jon",
		LastName:      "Jackson",
		VerifiedEmail: "jj@gmail.com",
		Password:      "jj",
	}

	_, err := rnuh.Handle(context.Background(), rnuc)
	if err != nil {
		t.Errorf("Find some errors %q", err)
	}

	stored := len(repo.Storage)
	fmt.Println("Stored users :", stored)
	if stored != 1 {
		t.Errorf("Number of stored users %v not equal to expected number %v", stored, 1)
	}

	cur := application.ConfirmUserRegistrationCommandHandler{Repository: repo}

	for _, u := range rnuh.Repository.(mockRepository).Storage {
		curc := application.ConfirmUserRegistrationCommand{ID: u.ID}

		_, err = cur.Handle(context.Background(), curc)
		if err != nil {
			t.Errorf("Find some errors %q", err)
		} else {
			println("Confirmed")
		}

	}
}
