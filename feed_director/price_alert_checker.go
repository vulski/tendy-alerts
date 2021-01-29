package feed_director

import (
	"github.com/vulski/tendy-alerts"
	"log"
)

type PriceChecker struct {
	notifierFactory tendy_alerts.NotifierFactory
	alertRepo       tendy_alerts.AlertRepository
	alertEval       tendy_alerts.AlertEvaluator
}

func NewPriceChecker(notifierFactory tendy_alerts.NotifierFactory, alertRepo tendy_alerts.AlertRepository) *PriceChecker {
	return &PriceChecker{
		notifierFactory: notifierFactory,
		alertRepo:       alertRepo,
	}
}

// Get active alerts for the given currency,
// Check alerts based on the latest prices,
// If Alert is Valid, notify user based on User settings.
func (p *PriceChecker) CheckPrice(price tendy_alerts.PriceSnapshot) error {
	// TODO: Optimize search for alerts based on the current price.
	alerts, err := p.alertRepo.GetActiveAlertsForCurrency(price.Currency)
	if err != nil {
		return handleErr(err)
	}

	for _, alert := range alerts {
		if p.alertEval.ShouldAlertUser(price, alert) {
			notifier, err := p.notifierFactory.CreateNotifierFromType(alert.NotificationSettings.Type)
			if err != nil {
				return handleErr(err)
			}
			err = notifier.NotifyUser(price, alert)
			if err != nil {
				return handleErr(err)
			}
			// TODO: update after adding logic for percentage alerts.
			alert.Active = false
			_, err = p.alertRepo.Save(alert)
			if err != nil {
				return handleErr(err)
			}
		}
	}

	return nil
}

func handleErr(err error) error {
	log.Printf("error: %s\n", err.Error())
	return err
}
