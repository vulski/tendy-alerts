package feed_director

import (
	ta "github.com/vulski/tendy-alerts"
	"log"
	"math"
	"time"
)

type PriceChecker interface {
	CheckPrice(price ta.PriceSnapshot) error
	LogPrice(price ta.PriceSnapshot) error
}

type PriceCheckerImpl struct {
	notifierFactory ta.NotifierFactory
	alertRepo       ta.AlertRepository
	priceRepo       ta.PriceSnapshotRepository
}

func NewPriceChecker(notifierFactory ta.NotifierFactory, alertRepo ta.AlertRepository, priceRepo ta.PriceSnapshotRepository) *PriceCheckerImpl {
	return &PriceCheckerImpl{
		notifierFactory: notifierFactory,
		alertRepo:       alertRepo,
		priceRepo:       priceRepo,
	}
}

func (a *PriceCheckerImpl) LogPrice(price ta.PriceSnapshot) error {
	previous, err := a.priceRepo.GetLatestForFrequency(ta.FifteenMinuteFrequency)
	if err != nil {
		return err
	}
	if previous.CreatedAt.Sub(price.CreatedAt).Minutes() > float64(time.Minute * 15){
		_, err = a.priceRepo.Save(price)
		if err != nil {
			return err
		}
	}

	return nil
}

// Get active alerts for the given currency,
// Check alerts based on the latest prices,
// If Alert is Valid, notify user based on User settings.
func (a *PriceCheckerImpl) CheckPrice(price ta.PriceSnapshot) error {
	// TODO: This is doing too many thing, maybe refactor later?
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

func (p *PriceCheckerImpl) shouldAlertUser(latestPrice ta.PriceSnapshot, alert ta.Alert) bool {
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
