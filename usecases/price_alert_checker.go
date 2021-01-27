package usecases

import (
	"github.com/vulski/tendy-alerts"
)

type PriceAlertChecker struct {
	notifierFactory tendy_alerts.NotifierFactory
	alertRepo       tendy_alerts.AlertRepository
	alertEval       tendy_alerts.AlertEvaluator
}

func NewPriceNotificationManager(notifierFactory tendy_alerts.NotifierFactory, alertRepo tendy_alerts.AlertRepository) *PriceAlertChecker {
	return &PriceAlertChecker{
		notifierFactory: notifierFactory,
		alertRepo:       alertRepo,
	}
}

// Get active alerts for the given currency,
// Check alerts based on the latest prices,
// If Alert is Valid, notify user based on User settings.
func (p *PriceAlertChecker) CheckPrice(price tendy_alerts.CurrencyPriceLog) error {
	alerts, err := p.alertRepo.GetActiveAlertsForCurrency(price.Currency)
	if err != nil {
		// TODO:
	}

	for _, alert := range alerts {
		if p.alertEval.ShouldAlertUser(price, alert) {
			notifier, err := p.notifierFactory.CreateNotifierFromType(alert.NotificationSettings.Type)
			if err != nil {
				panic(err)
			}
			err = notifier.NotifyUser(price, alert)
			if err != nil {
				// TODO:
			}
		}
	}

	return nil
}
