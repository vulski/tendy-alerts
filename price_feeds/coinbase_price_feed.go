package price_feeds

import (
	ws "github.com/gorilla/websocket"
	"github.com/preichenberger/go-coinbasepro/v2"
	tendy_alerts "github.com/vulski/tendy-alerts"
	"strconv"
	"strings"
	"time"
)

type CoinBasePriceFeed struct {
	wsConn            *ws.Conn
	watchedCurrencies map[string]chan tendy_alerts.PriceSnapshot
	running           bool
}

func NewCoinBasePriceFeed() *CoinBasePriceFeed {
	return &CoinBasePriceFeed{watchedCurrencies: make(map[string]chan tendy_alerts.PriceSnapshot)}
}

func (c *CoinBasePriceFeed) ExchangeName() string {
	return "coinbase"
}

func (c *CoinBasePriceFeed) SubscribeToCurrency(currency string) (chan tendy_alerts.PriceSnapshot, error) {
	err := c.subscribeToCoin(currency)
	if err != nil {
		return nil, err
	}

	return c.watchedCurrencies[currency], nil
}

func (c *CoinBasePriceFeed) StartFeed() error {
	c.running = true
	go c.digestCoins()
	return nil
}

func (c *CoinBasePriceFeed) StopFeed() error {
	c.running = false
	err := c.wsConn.Close()
	if err != nil {
		return err
	}
	return nil
}

func (c *CoinBasePriceFeed) subscribeToCoin(coin string) error {
	var wsDialer ws.Dialer
	wsConn, _, err := wsDialer.Dial("wss://ws-feed.pro.coinbase.com", nil)
	if err != nil {
		return err
	}
	c.wsConn = wsConn

	subscribe := coinbasepro.Message{
		Type: "subscribe",
		Channels: []coinbasepro.MessageChannel{
			{
				Name: "ticker",
				ProductIds: []string{
					coin + "-USD",
				},
			},
		},
	}
	if err := c.wsConn.WriteJSON(subscribe); err != nil {
		return err
	}
	c.watchedCurrencies[coin] = make(chan tendy_alerts.PriceSnapshot)

	return nil
}

func (c *CoinBasePriceFeed) digestCoins() {
	for c.running {
		message := coinbasepro.Message{}
		if err := c.wsConn.ReadJSON(&message); err != nil {
			println(err.Error())
			break
		}
		price, err := strconv.ParseFloat(message.BestBid, 64)
		if nil != err {
			continue
		}
		currency := strings.Split(message.ProductID, "-")
		if len(currency) > 0 {
			currency := currency[0]
			latestPrice := tendy_alerts.PriceSnapshot{
				Price:     price,
				Currency:  currency,
				Exchange:  c.ExchangeName(),
				CreatedAt: time.Now(),
			}
			c.watchedCurrencies[currency] <- latestPrice
		}
	}
}
