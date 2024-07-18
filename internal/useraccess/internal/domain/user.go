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

func (u User) GetID() UserID {
	return u.id
}

func (u User) GetEmail() UserEmail {
	return u.email
}

func (u User) GetUserName() UserName {
	return u.userName
}

func (u User) GetPassword() HashedUserPassword {
	return u.password
}
