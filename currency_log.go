package tendy_alerts

import "time"

type CurrencyPriceLog struct {
	Currency  string  `json:"stub"`
	Price     float64 `json:"price"`
	Exchange  string  `json:"exchange"`
	Timestamp time.Time
}

type CurrencyPriceLogRepoConstraints struct {
	Name             string
	PriceLessThan    string
	PriceGreaterThan string
}

type CurrencyPriceLogRepository interface {
	GetWithConstraints(constraints CurrencyPriceLogRepoConstraints) ([]*CurrencyPriceLog, error)
}
