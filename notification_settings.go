package tendy_alerts

import (
	"github.com/vulski/tendy-alerts/enums"
)

type NotificationSetting struct {
	Entity
	Type enums.NotificationType
	UserId uint
	AlertId uint
}

//go:generate mockgen -destination=mocks/mock_notification_setting_repository.go -package=mocks . NotificationSettingRepository
type NotificationSettingRepository interface {
	GetForAlertId(id uint) (NotificationSetting, error)
}
