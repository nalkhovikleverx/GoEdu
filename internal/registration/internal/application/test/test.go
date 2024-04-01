package test

import (
	"context"
	"errors"

	"GoEdu/internal/registration/internal/application"
	"GoEdu/internal/registration/internal/domain"
)

func NewMockRepository() *MockRepository {
	return &MockRepository{Storage: make([]domain.UserRegistration, 0)}
}

var _ application.UserRegistrationRepository = (*MockRepository)(nil)

type MockRepository struct {
	Storage []domain.UserRegistration
}

func (m *MockRepository) Add(ctx context.Context, registration *domain.UserRegistration) error {
	m.Storage = append(m.Storage, *registration)
	return nil
}

func (m *MockRepository) Load(ctx context.Context, urid domain.UserRegistrationID) (domain.UserRegistration, error) {
	for _, user := range m.Storage {
		if user.ID == urid {
			return user, nil
		}
	}
	return domain.UserRegistration{}, errors.New("user not found")
}

func (m *MockRepository) Update(ctx context.Context, user *domain.UserRegistration) error {
	for i := 0; i < len(m.Storage); i++ {
		if m.Storage[i].ID == user.ID {
			m.Storage[i] = *user
			return nil
		}
	}
	return errors.New("user not found")
}

func (m *MockRepository) GetByEmail(ctx context.Context, email domain.UserRegistrationEmail) (domain.UserRegistration, error) {
	for _, user := range m.Storage {
		if user.Email == email {
			return user, nil
		}
	}
	return domain.UserRegistration{}, errors.New("user not found")
}

type mockHasher struct {
}

func (m *mockHasher) Hash(password string) (string, error) {
	return password + "aaa", nil
}

type UniqueMailVerifier struct {
	Repo MockRepository
}

func (u *UniqueMailVerifier) IsUnique(ctx context.Context, email domain.UserRegistrationEmail) error {
	_, err := u.Repo.GetByEmail(ctx, email)
	if err != nil {
		return nil
	}
	return application.UserEmailMustBeUniqueError
}
