package models

import "time"

// TODO: how tf do you organize enums cleanly in Go?

type AlertType string
const (
	TargetAlert AlertType = "target_alert"
	PercentageChangeAlert = "percentage_change_alert"
)

type AlertFrequency string
const (
	OneTime AlertFrequency = "one_time"
	FifteenMinutes = "15m"
	ThirtyMinutes = "30m"
	OneHour = "1h"
	SixHours = "6h"
	TwelveHours = "12h"
	TwentyFourHours = "24h"
)

type Alert struct {
	Currency string
	Price string
	PercentageChange string
	Type AlertType
	Frequency AlertFrequency
	TradePair string
	Timestamp time.Time
}