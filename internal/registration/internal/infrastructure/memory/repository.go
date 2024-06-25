package memory

import (
	"context"
	"errors"

	"GoEdu/internal/registration/internal/application"
	"GoEdu/internal/registration/internal/domain"
)

var _ application.UserRegistrationRepository = (*InProcessRepository)(nil)

type InProcessRepository struct {
	store map[domain.UserRegistrationID]domain.UserRegistrationSnapshot
}

func NewRepository() *InProcessRepository {
	return &InProcessRepository{}
}

func (i *InProcessRepository) Add(_ context.Context, registration *domain.UserRegistration) error {
	snapShot := registration.GetSnapshot()
	i.store[snapShot.ID] = snapShot
	return nil
}

func (i *InProcessRepository) Load(
	_ context.Context,
	id domain.UserRegistrationID,
) (*domain.UserRegistration, error) {
	usr, ok := i.store[id]
	if ok {
		return domain.MustCreateUserRegistrationFromSnapshot(usr), nil
	}
	return nil, errors.New("user not found")
}

func (i *InProcessRepository) Update(_ context.Context, registration *domain.UserRegistration) error {
	snapShot := registration.GetSnapshot()
	if _, ok := i.store[snapShot.ID]; ok {
		i.store[snapShot.ID] = snapShot
		return nil
	}
	return errors.New("this user doesn't exist")
}
