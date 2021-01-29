package tendy_alerts

import "time"

type AlertComparison string
type AlertFrequency string
type AlertType string

const (
	GreaterThanComparison AlertComparison = "greater_than"
	LessThanComparison    AlertComparison = "less_than"

	FifteenMinuteFrequency  AlertFrequency = "15m"
	ThirtyMinuteFrequency                  = "30m"
	OneHourFrequency                       = "1hr"
	SixHourFrequency                       = "6hr"
	TwelveHourFrequency                    = "12hr"
	TwentyFourHourFrequency                = "24hr"

	TargetAlert           AlertType = "target_alert"
	PercentageChangeAlert           = "percentage_change_alert"
)

type Alert struct {
	ID                   uint
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            time.Time
	Currency             string
	Price                float64
	PercentageChange     float64
	Type                 AlertType
	Frequency            AlertFrequency
	Comparison           AlertComparison
	TradePair            string
	Active               bool
	UserId               uint
	NotificationSettings NotificationSettings
}

//go:generate mockgen -destination=mocks/mock_alert_repository.go -package=mocks . AlertRepository
type AlertRepository interface {
	Save(alert Alert) (*Alert, error)
	GetActiveAlertsForCurrency(currency string) ([]Alert, error)
}

type AlertRepositoryInMem struct {
	Alerts []Alert
}

func (r *AlertRepositoryInMem) Save(alert Alert) (*Alert, error) {
	r.Alerts = []Alert{alert}
	return &alert, nil
}

func (r *AlertRepositoryInMem) GetActiveAlertsForCurrency(currency string) ([]Alert, error) {
	return r.Alerts, nil
}
