package tendy_alerts

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

type NotificationSetting struct {
	Entity
	Type    NotificationType
	UserId  uint
	AlertId uint
}

//go:generate mockgen -destination=mocks/mock_notification_setting_repository.go -package=mocks . NotificationSettingRepository
type NotificationSettingRepository interface {
	GetForAlertId(id uint) (NotificationSetting, error)
}
