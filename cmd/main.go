package main

import (
	"github.com/vulski/tendy-alerts"
	"github.com/vulski/tendy-alerts/feed_director"
	"github.com/vulski/tendy-alerts/notifiers"
)

func main() {
	targetAlert := tendy_alerts.Alert{
		Currency:             "BTC",
		Price:                34_700,
		Type:                 tendy_alerts.TargetAlert,
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
	priceChecker := feed_director.NewPriceChecker(factory, &alertRepo, &tendy_alerts.PriceSnapshotRepoInMem{})
	director := feed_director.New(priceChecker)
	director.Start()
	for {
		// TODO: CLI or something
	}
}
