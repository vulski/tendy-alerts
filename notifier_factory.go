package tendy_alerts

import "github.com/vulski/tendy-alerts/enums"

//go:generate mockgen -destination=mocks/mock_notifier_factory.go -package=mocks . NotifierFactory
type NotifierFactory interface {
	CreateNotifierFromType(notificationType enums.NotificationType) (Notifier, error)
}
