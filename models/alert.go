package models

import (
	"github.com/vulski/tendy-alerts/enums"
	"time"
)

type Alert struct {
	Currency         string
	Price            float64
	PercentageChange float64
	Type             enums.AlertType
	Frequency        enums.AlertFrequency
	Comparison       enums.AlertComparison
	TradePair        string
	Timestamp        time.Time
}
