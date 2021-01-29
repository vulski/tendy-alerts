package feed_director

import (
	ta "github.com/vulski/tendy-alerts"
	"log"
	"math"
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
func (a *PriceChecker) CheckPrice(price ta.PriceSnapshot) error {
	// TODO: Optimize search for alerts based on the current price.
	alerts, err := a.alertRepo.GetActiveAlertsForCurrency(price.Currency)
	if err != nil {
		return handleErr(err)
	}
	for _, alert := range alerts {
		if a.shouldAlertUser(price, alert) {
			notifier, err := a.notifierFactory.CreateNotifierFromType(alert.NotificationSettings.Type)
			if err != nil {
				return handleErr(err)
			}
			err = notifier.NotifyUser(price, alert)
			if err != nil {
				return handleErr(err)
			}
			if alert.Type == ta.TargetAlert {
				alert.Active = false
				_, err = a.alertRepo.Save(alert)
				if err != nil {
					return handleErr(err)
				}
			}
		}
	}

	return nil
}

func (p *PriceChecker) shouldAlertUser(latestPrice ta.PriceSnapshot, alert ta.Alert) bool {
	if !alert.Active {
		return false
	}
	switch alert.Type {
	case ta.TargetAlert:
		if alert.Comparison == ta.LessThanComparison {
			return latestPrice.Price < alert.Price
		}
		return latestPrice.Price > alert.Price

	case ta.PercentageChangeAlert:
		price, err := p.priceRepo.GetLatestForFrequency(alert.Frequency)
		if err != nil {
			// TODO: Maybe have some retry attempts?
			log.Println("error: " + err.Error())
			return false
		}
		change := latestPrice.Price - price.Price
		percentageChange := change / price.Price
		return math.Abs(percentageChange) > alert.PercentageChange
	default:
		return false
	}
}

func handleErr(err error) error {
	log.Printf("error: %s\n", err.Error())
	return err
}
