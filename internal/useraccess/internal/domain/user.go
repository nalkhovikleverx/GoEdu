package domain

type User struct {
	id       UserID
	email    UserEmail
	userName UserName
	password HashedUserPassword
}

func NewUser(
	userName UserName,
	password string,
	email UserEmail,
) (*User, error) {
	return &User{
		NewUserID(),
		email,
		userName,
		NewHashedUserPassword(MustNewUserPassword(password)),
	}, nil
}

func (u *User) GetUserSnapshot() *UserSnapshot {
	return &UserSnapshot{
		ID:       u.id,
		Email:    u.email,
		UserName: u.userName,
		Password: u.password,
	}
}

func (u *User) IsPasswordEqual(password UserPassword) bool {
	return NewHashedUserPassword(password) == u.password
}

type UserSnapshot struct {
	ID       UserID
	Email    UserEmail
	UserName UserName
	Password HashedUserPassword
}
