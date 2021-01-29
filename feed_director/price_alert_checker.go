package feed_director

import (
	ta "github.com/vulski/tendy-alerts"
	"log"
)

type PriceChecker struct {
	notifierFactory ta.NotifierFactory
	alertRepo       ta.AlertRepository
	priceRepo       ta.PriceSnapshotRepository
}

func NewPriceChecker(notifierFactory ta.NotifierFactory, alertRepo ta.AlertRepository, priceRepo ta.PriceSnapshotRepository) *PriceChecker {
	return &PriceChecker{
		notifierFactory: notifierFactory,
		alertRepo:       alertRepo,
		priceRepo:       priceRepo,
	}
}

// Get active alerts for the given currency,
// Check alerts based on the latest prices,
// If Alert is Valid, notify user based on User settings.
func (p *PriceChecker) CheckPrice(price ta.PriceSnapshot) error {
	// TODO: Optimize search for alerts based on the current price.
	alerts, err := p.alertRepo.GetActiveAlertsForCurrency(price.Currency)
	if err != nil {
		return handleErr(err)
	}
	for _, alert := range alerts {
		if p.shouldAlertUser(price, alert) {
			notifier, err := p.notifierFactory.CreateNotifierFromType(alert.NotificationSettings.Type)
			if err != nil {
				return handleErr(err)
			}
			err = notifier.NotifyUser(price, alert)
			if err != nil {
				return handleErr(err)
			}
			if alert.Type == ta.TargetAlert {
				alert.Active = false
				_, err = p.alertRepo.Save(alert)
				if err != nil {
					return handleErr(err)
				}
			}
		}
	}

	return nil
}

func (a *PriceChecker) shouldAlertUser(latestPrice ta.PriceSnapshot, alert ta.Alert) bool {
	if !alert.Active {
		return false
	}

	if alert.Type == ta.TargetAlert {
		if alert.Comparison == ta.LessThanComparison {
			return latestPrice.Price < alert.Price
		}
		return latestPrice.Price > alert.Price
	}
	return false
}

func handleErr(err error) error {
	log.Printf("error: %s\n", err.Error())
	return err
}
