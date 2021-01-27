package main

import (
	"fmt"
	ws "github.com/gorilla/websocket"
	"github.com/vulski/tendy-alerts"
	"github.com/vulski/tendy-alerts/manager"
	"github.com/vulski/tendy-alerts/notifiers"
	"github.com/vulski/tendy-alerts/price_feeds"
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
	coinbaseFeed := price_feeds.NewCoinBasePriceFeed()
	btcChan, err := coinbaseFeed.GetCurrencyFeed(targetAlert.Currency)
	alertRepo := tendy_alerts.AlertRepositoryInMem{Alerts: []tendy_alerts.Alert{targetAlert}}
	priceChecker := manager.NewPriceAlertChecker(notifiers.NewNotifierFactory(), &alertRepo)
	if err != nil {
		fmt.Println(err)
		return
	}
	if btcChan == nil {
		panic("ahh")
		return
	}
	coinbaseFeed.StartFeed()
	for {
		select {
		case latestPrice := <-btcChan:
			priceChecker.CheckPrice(latestPrice)
			fmt.Println(latestPrice)
		}
	}
}
