package tendy_alerts

type User struct {
	Entity
	Username string
}

type UserRepository interface {
	GetAllActive() ([]*User, error)
}

type UserRepoMock struct {
	Users []*User
}

func (u *UserRepoMock) GetAllActive() ([]*User, error) {
	return u.Users, nil
}
