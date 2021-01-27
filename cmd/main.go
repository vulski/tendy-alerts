package main

import (
	"fmt"
	ws "github.com/gorilla/websocket"
	"github.com/preichenberger/go-coinbasepro/v2"
	"github.com/vulski/tendy-alerts"
	"github.com/vulski/tendy-alerts/enums"
	"strconv"
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
		Comparison: enums.AlertComparison_LESS_THAN,
		TradePair:  "BTC/USD",
		Active:     true,
	}
	alerts = append(alerts, &targetAlert)
	var wsDialer ws.Dialer
	wsConn, _, err := wsDialer.Dial("wss://ws-feed.pro.coinbase.com", nil)
	if err != nil {
		println(err.Error())
	}

	subscribe := coinbasepro.Message{
		Type: "subscribe",
		Channels: []coinbasepro.MessageChannel{
			{
				Name: "ticker",
				ProductIds: []string{
					"BTC-USD",
				},
			},
		},
	}
	if err := wsConn.WriteJSON(subscribe); err != nil {
		println(err.Error())
	}
	for {
		message := coinbasepro.Message{}
		if err := wsConn.ReadJSON(&message); err != nil {
			println(err.Error())
			break
		}
		price, err := strconv.ParseFloat(message.BestBid, 64)
		if nil != err {
			continue
		}
		latestPrice := tendy_alerts.CurrencyPriceLog{Price: price}
		fmt.Println(latestPrice.Price)

		for _, alert := range alerts {
			fmt.Println("Evaluating ")
			if alertEval.ShouldAlertUser(latestPrice, *alert) {
				fmt.Println("BUYBUYBUY")
				alert.Active = false
			}
		}
	}
}
