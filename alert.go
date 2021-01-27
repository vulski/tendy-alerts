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
}

type AlertRepository interface {
	GetAllForUserIDs(userIds []uint) ([]Alert, error)
}

type AlertRepoMock struct {
	Alerts []Alert
}

func (u *AlertRepoMock) GetAllForUserIDs(ids uint) ([]Alert, error) {
	return u.Alerts, nil
}
