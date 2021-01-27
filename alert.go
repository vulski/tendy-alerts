package tendy_alerts

import (
	"github.com/vulski/tendy-alerts/enums"
)

type Alert struct {
	Entity
	Currency             string
	Price                float64
	PercentageChange     float64
	Type                 enums.AlertType
	Frequency            enums.AlertFrequency
	Comparison           enums.AlertComparison
	TradePair            string
	Active               bool
	UserId               uint
	NotificationSettings NotificationSetting
}

//go:generate mockgen -destination=mocks/mock_alert_repository.go -package=mocks . AlertRepository
type AlertRepository interface {
	GetActiveAlertsForCurrency(currency string) ([]Alert, error)
}

type AlertRepositoryInMem struct {
	Alerts []Alert
}

func (r *AlertRepositoryInMem) GetActiveAlertsForCurrency(currency string) ([]Alert, error) {
	return r.Alerts, nil
}
