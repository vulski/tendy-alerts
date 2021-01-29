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

//go:generate mockgen -destination=mocks/mock_price_snapshot_repository.go -package=mocks . PriceSnapshotRepository
type PriceSnapshotRepository interface {
	GetLatestForFrequency(freq AlertFrequency) (PriceSnapshot, error)
}
