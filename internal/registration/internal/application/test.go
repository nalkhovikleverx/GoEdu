package application

import (
	"context"

	"GoEdu/internal/registration/internal/domain"
)

var _ UserRegistrationRepository = (*UserRepositorySpy)(nil)

type UserRepositorySpy struct {
	Added   *domain.UserRegistration
	Loaded  *domain.UserRegistration
	Updated *domain.UserRegistration
}

func NewUserRepositorySpy(Added *domain.UserRegistration, Loaded *domain.UserRegistration, Updated *domain.UserRegistration) *UserRepositorySpy {
	return &UserRepositorySpy{Added: Added, Loaded: Loaded, Updated: Updated}
}

func (u *UserRepositorySpy) Add(_ context.Context, user *domain.UserRegistration) error {
	u.Added = user
	return nil
}

func (u *UserRepositorySpy) Load(_ context.Context, _ domain.UserRegistrationID) (*domain.UserRegistration, error) {
	return u.Loaded, nil
}

func (u *UserRepositorySpy) Update(_ context.Context, user *domain.UserRegistration) error {
	u.Updated = user
	return nil
}

var _ UniqueEmailVerifier = (*UniqueEmailVerifierSpy)(nil)

type UniqueEmailVerifierSpy struct {
	checked *domain.UserRegistrationEmail
}

func (u *UniqueEmailVerifierSpy) IsUnique(_ context.Context, email *domain.UserRegistrationEmail) error {
	u.checked = email
	return nil
}

var _ PasswordHasher = (*PasswordHasherSpy)(nil)

type PasswordHasherSpy struct {
	hashed *domain.UserPassword
}

func (p *PasswordHasherSpy) Hash(password *domain.UserPassword) (*domain.UserPassword, error) {
	p.hashed = password
	return password, nil
}
