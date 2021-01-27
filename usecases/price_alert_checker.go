package usecases

import (
	"github.com/vulski/tendy-alerts"
)

// Get users
// Get each users active alerts
// Pull in exchange price feeds
// Check User's Alerts based on the latest prices
// If Alert is Valid, notify user based on User settings.

// Update this User and User Alert list regularly

type PriceAlertChecker struct {
	userRepository          tendy_alerts.UserRepository
	alertRepository         tendy_alerts.AlertRepository
	notificationSettingRepo tendy_alerts.NotificationSettingRepository
	notifier                tendy_alerts.Notifier

	users                   []*tendy_alerts.User
	alertEval               tendy_alerts.AlertEvaluator
}

func NewPriceNotificationManager(userRepo tendy_alerts.UserRepository) *PriceAlertChecker {
	return &PriceAlertChecker{userRepository: userRepo}
}

func (p *PriceAlertChecker) Run() error {
	users, err := p.userRepository.GetAllActive()
	if nil != err {
		// TODO:
	}
	var userIds []uint
	for _, user := range users {
		userIds = append(userIds, user.ID)
	}

	alerts, err := p.alertRepository.GetAllForUserIDs(userIds)
	if err != nil {
		return err
	}

	latestPrice := tendy_alerts.CurrencyPriceLog{}
	for _, alert := range alerts {
		if p.alertEval.ShouldAlertUser(latestPrice, alert) {
			notificationSetting, err := p.notificationSettingRepo.GetForAlertId(alert.ID)
			if err != nil {
				// TODO:
			}
			err = p.notifier.NotifyUser(latestPrice, alert, notificationSetting)
			if err != nil {
				// TODO:
			}
		}
	}

	return nil
}
