package main

import (
	"github.com/vulski/tendy-alerts"
	"github.com/vulski/tendy-alerts/manager"
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
		NotificationSettings: tendy_alerts.NotificationSettings{Type: tendy_alerts.EmailNotification},
	}
	factory, err := notifiers.NewNotifierFactoryFromConfig("config/notifiers.json")
	if err != nil {
		panic(err)
	}
	alertRepo := tendy_alerts.AlertRepositoryInMem{Alerts: []tendy_alerts.Alert{targetAlert}}
	priceChecker := manager.NewPriceChecker(factory, &alertRepo)
	mngr := manager.New(priceChecker)
	mngr.Start()
	for {
		// TODO: CLI or something
	}
}
