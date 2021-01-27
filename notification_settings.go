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

type NotificationSettingRepository interface {
	GetForAlertId(id uint) (NotificationSetting, error)
}

type NotificationSettingRepoMock struct {
	NotSetting NotificationSetting
}

func (u *NotificationSettingRepoMock) GetForAlertId(id uint) (NotificationSetting, error) {
	return u.NotSetting, nil
}
