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
	notifierFactory tendy_alerts.NotifierFactory
	userRepository  tendy_alerts.UserRepository
	alertRepository tendy_alerts.AlertRepository
	notifSettRepo   tendy_alerts.NotificationSettingRepository

	alertEval tendy_alerts.AlertEvaluator
}

func NewPriceNotificationManager(notifierFactory tendy_alerts.NotifierFactory, userRepo tendy_alerts.UserRepository, alertRepo tendy_alerts.AlertRepository, notifiSettRepo tendy_alerts.NotificationSettingRepository) *PriceAlertChecker {
	return &PriceAlertChecker{
		notifierFactory: notifierFactory,
		userRepository:  userRepo,
		alertRepository: alertRepo,
		notifSettRepo:   notifiSettRepo,
	}
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

	// TODO: PriceFetcher
	latestPrice := tendy_alerts.CurrencyPriceLog{}

	for _, alert := range alerts {
		if p.alertEval.ShouldAlertUser(latestPrice, alert) {
			notificationSettings, err := p.notifSettRepo.GetForAlertId(alert.ID)
			if nil != err {
				// TODO:
			}
			notifier, err := p.notifierFactory.CreateNotifierFromType(notificationSettings.Type)
			err = notifier.NotifyUser(latestPrice, alert)
			if err != nil {
				// TODO:
			}
		}
	}

	return nil
}
