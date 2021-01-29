package tendy_alerts

import (
	"fmt"
	"time"
)

type PriceFeed interface {
	ExchangeName() string
	SubscribeToCurrency(currency string) (chan PriceSnapshot, error)
	StartFeed() error
	StopFeed() error
}
type PriceSnapshot struct {
	Currency  string  `json:"stub"`
	Price     float64 `json:"price"`
	Exchange  string  `json:"exchange"`
	Timestamp time.Time
}

// TODO: make fancy.
func (p *PriceSnapshot) Stringify() string {
	return fmt.Sprintf("%s's price has changed to %f", p.Currency, p.Price)
}

type CurrencyPriceLogRepoConstraints struct {
	Name             string
	PriceLessThan    string
	PriceGreaterThan string
}

type CurrencyPriceLogRepository interface {
	GetWithConstraints(constraints CurrencyPriceLogRepoConstraints) ([]*PriceSnapshot, error)
}
