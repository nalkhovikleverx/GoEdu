package inprocess

import (
	"context"
	"errors"

	"GoEdu/internal/registration/internal/application"
	"GoEdu/internal/registration/internal/domain"
)

var _ application.UserRegistrationRepository = (*InProcessRepository)(nil)

type InProcessRepository struct {
	store []domain.UserRegistration
}

func NewRepository() *InProcessRepository {
	return &InProcessRepository{}
}

func (i *InProcessRepository) Add(ctx context.Context, registration *domain.UserRegistration) error {
	i.store = append(i.store, *registration)
	return nil
}

func (i *InProcessRepository) Load(
	ctx context.Context,
	id domain.UserRegistrationID,
) (*domain.UserRegistration, error) {
	for _, user := range i.store {
		if user.GetSnapshot().ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (i *InProcessRepository) Update(ctx context.Context, registration *domain.UserRegistration) error {
	for ind := 0; ind < len(i.store); ind++ {
		val := i.store[ind]
		if val.GetSnapshot().ID == registration.GetSnapshot().ID {
			i.store[ind] = *registration
			return nil
		}
	}
	return nil
}
