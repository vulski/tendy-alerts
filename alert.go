package tendy_alerts

import (
	"github.com/vulski/tendy-alerts/enums"
)

type Alert struct {
	Entity
	Currency         string
	Price            float64
	PercentageChange float64
	Type             enums.AlertType
	Frequency        enums.AlertFrequency
	Comparison       enums.AlertComparison
	TradePair        string
	Active           bool
	UserId           uint
}

//go:generate mockgen -destination=mocks/mock_alert_repository.go -package=mocks . AlertRepository
type AlertRepository interface {
	GetAllForUserIDs(userIds []uint) ([]Alert, error)
}