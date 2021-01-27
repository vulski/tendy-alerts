package main

import (
	"fmt"
	ws "github.com/gorilla/websocket"
	"github.com/vulski/tendy-alerts"
	"github.com/vulski/tendy-alerts/enums"
	"github.com/vulski/tendy-alerts/usecases"
	"github.com/vulski/tendy-alerts/usecases/price_feeds"
)

var alerts []*tendy_alerts.Alert
var alertEval *tendy_alerts.AlertEvaluator
var wsConn *ws.Conn

func init() {
	alertEval = tendy_alerts.NewAlertEvaluator()
}

func main() {
	targetAlert := tendy_alerts.Alert{
		Currency:   "BTC",
		Price:      20000,
		Type:       enums.AlertType_TARGET_ALERT,
		Frequency:  enums.AlertFrequency_ONE_TIME,
		Comparison: enums.AlertComparison_GREATER_THAN,
		TradePair:  "BTC/USD",
		Active:     true,
		NotificationSettings: tendy_alerts.NotificationSetting{Type: enums.NotificationType_EMAIL},
	}
	coinbaseFeed := price_feeds.NewCoinBasePriceFeed()
	btcChan, err := coinbaseFeed.GetCurrencyFeed(targetAlert.Currency)
	alertRepo := tendy_alerts.AlertRepositoryInMem{Alerts: []tendy_alerts.Alert{targetAlert}}
	priceChecker := usecases.NewPriceNotificationManager(usecases.NewNotifierFactory(), &alertRepo)
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
