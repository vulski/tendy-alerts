package usecases

import (
	"github.com/vulski/tendy-alerts"
)

type PriceAlertChecker struct {
	notifierFactory tendy_alerts.NotifierFactory
	userRepository  tendy_alerts.UserRepository
	alertEval tendy_alerts.AlertEvaluator
}

func NewPriceNotificationManager(notifierFactory tendy_alerts.NotifierFactory, userRepo tendy_alerts.UserRepository) *PriceAlertChecker {
	return &PriceAlertChecker{
		notifierFactory: notifierFactory,
		userRepository:  userRepo,
	}
}

// Get users
// Get each users active alerts
// Pull in exchange price feeds
// Check User's Alerts based on the latest prices
// If Alert is Valid, notify user based on User settings.
func (p *PriceAlertChecker) Run() error {
	users, err := p.userRepository.GetAllActiveWithAlerts()
	if nil != err {
		// TODO:
	}

	// TODO: PriceFetcher
	latestPrice := tendy_alerts.CurrencyPriceLog{}

	for _, user := range users {
		for _, alert := range user.Alerts {
			if p.alertEval.ShouldAlertUser(latestPrice, alert) {
				notifier, err := p.notifierFactory.CreateNotifierFromType(alert.NotificationSettings.Type)
				err = notifier.NotifyUser(latestPrice, alert)
				if err != nil {
					// TODO:
				}
			}
		}
	}

	return nil
}
