package domain

import "github.com/google/uuid"

type UserRegistrationID uuid.UUID

func NewUserRegistrationID() UserRegistrationID {
	return UserRegistrationID(uuid.New())
}

func ParseUserRegistrationID(urid string) (UserRegistrationID, error) {
	uid, err := uuid.Parse(urid)
	if err != nil {
		return UserRegistrationID{}, err
	}
	return UserRegistrationID(uid), nil
}

func MustParseUserRegistrationID(urid string) UserRegistrationID {
	uid, err := ParseUserRegistrationID(urid)
	if err != nil {
		panic(err)
	}
	return uid
}
