package main

import (
	"github.com/vulski/tendy-alerts"
	"github.com/vulski/tendy-alerts/feed_director"
	"github.com/vulski/tendy-alerts/notifiers"
)

func main() {
	targetAlert := tendy_alerts.Alert{
		Currency:             "BTC",
		Price:                20000,
		Type:                 tendy_alerts.TargetAlert,
		Frequency:            tendy_alerts.OneTimeFrequency,
		Comparison:           tendy_alerts.GreaterThanComparison,
		TradePair:            "BTC/USD",
		Active:               true,
		NotificationSettings: tendy_alerts.NotificationSettings{Type: tendy_alerts.EmailNotification, TargetUsername: "to@example.com"},
	}
	factory, err := notifiers.NewNotifierFactoryFromConfig("config/notifiers.json")
	if err != nil {
		panic(err)
	}
	alertRepo := tendy_alerts.AlertRepositoryInMem{Alerts: []tendy_alerts.Alert{targetAlert}}
	priceChecker := feed_director.NewPriceChecker(factory, &alertRepo)
	mngr := feed_director.New(priceChecker)
	mngr.Start()
	for {
		// TODO: CLI or something
	}
}
