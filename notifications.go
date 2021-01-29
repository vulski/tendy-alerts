package tendy_alerts

import (
	"time"
)

//go:generate mockgen -destination=mocks/mock_notifier.go -package=mocks . Notifier
type Notifier interface {
	NotifyUser(currencyLog PriceSnapshot, alert Alert) error
}

//go:generate mockgen -destination=mocks/mock_notifier_factory.go -package=mocks . NotifierFactory
type NotifierFactory interface {
	CreateNotifierFromType(notificationType NotificationType) (Notifier, error)
}

type NotificationType string

const EmailNotification NotificationType = "email"

type NotificationSettings struct {
	ID             uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
	Type           NotificationType
	TargetUsername string
	UserId         uint
	AlertId        uint
}

//go:generate mockgen -destination=mocks/mock_notification_setting_repository.go -package=mocks . NotificationSettingsRepository
type NotificationSettingsRepository interface {
	GetForAlertId(id uint) (NotificationSettings, error)
}
