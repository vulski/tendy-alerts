package models

import "time"

type CurrencyPriceLog struct {
	Stub      string `json:"stub"`
	Price     string `json:"price"`
	Timestamp time.Time
}

type CurrencyPriceLogRepoConstraints struct {
	Name string
	PriceLessThan string
	PriceGreaterThan string
}

type CurrencyPriceLogRepository interface {
	GetWithConstraints(constraints CurrencyPriceLogRepoConstraints) ([]*CurrencyPriceLog, error)
}