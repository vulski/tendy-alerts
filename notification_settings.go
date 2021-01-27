package tendy_alerts

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
