package tendy_alerts

type User struct {
	Entity
	Username string
}

//go:generate mockgen -destination=mocks/mock_user_repository.go -package=mocks . UserRepository
type UserRepository interface {
	GetAllActive() ([]*User, error)
}

