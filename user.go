package tendy_alerts

type User struct {
	Entity
	Username string
	Alerts   []Alert
}

//go:generate mockgen -destination=mocks/mock_user_repository.go -package=mocks . UserRepository
type UserRepository interface {
	GetAllActiveWithAlerts() ([]*User, error)
}
