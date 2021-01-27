package main

import (
	ws "github.com/gorilla/websocket"
	"github.com/vulski/tendy-alerts"
	"github.com/vulski/tendy-alerts/manager"
	"github.com/vulski/tendy-alerts/notifiers"
)

var alerts []*tendy_alerts.Alert
var alertEval *tendy_alerts.AlertEvaluator
var wsConn *ws.Conn

func init() {
	alertEval = tendy_alerts.NewAlertEvaluator()
}

func main() {
	targetAlert := tendy_alerts.Alert{
		Currency:             "BTC",
		Price:                20000,
		Type:                 tendy_alerts.TargetAlert,
		Frequency:            tendy_alerts.OneTimeFrequency,
		Comparison:           tendy_alerts.GreaterThanComparison,
		TradePair:            "BTC/USD",
		Active:               true,
		NotificationSettings: tendy_alerts.NotificationSetting{Type: tendy_alerts.EmailNotification},
	}
	alertRepo := tendy_alerts.AlertRepositoryInMem{Alerts: []tendy_alerts.Alert{targetAlert}}
	priceChecker := manager.NewPriceChecker(notifiers.NewNotifierFactory(), &alertRepo)
	mngr := manager.New(priceChecker)
	mngr.Start()
	for {
		// TODO: CLI or something
	}
}
