package tendy_alerts

type AlertComparison string
type AlertFrequency string
type AlertType string

const (
	GreaterThanComparison AlertComparison = "greater_than"
	LessThanComparison    AlertComparison = "less_than"

	OneTimeFrequency        AlertFrequency = "one_time"
	FifteenMinuteFrequency                 = "15m"
	ThirtyMinuteFrequency                  = "30m"
	OneHourFrequency                       = "1hr"
	SixHourFrequency                       = "6hr"
	TwelveHourFrequency                    = "12hr"
	TwentyFourHourFrequency                = "24hr"

	TargetAlert           AlertType = "target_alert"
	PercentageChangeAlert           = "percentage_change_alert"
)

type Alert struct {
	Entity
	Currency             string
	Price                float64
	PercentageChange     float64
	Type                 AlertType
	Frequency            AlertFrequency
	Comparison           AlertComparison
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

type AlertEvaluator struct {
}

func NewAlertEvaluator() *AlertEvaluator {
	return &AlertEvaluator{}
}

func (a *AlertEvaluator) ShouldAlertUser(latestPrice CurrencyPriceLog, alert Alert) bool {
	if !alert.Active {
		return false
	}

	if alert.Type == TargetAlert {
		if alert.Comparison == LessThanComparison {
			return latestPrice.Price < alert.Price
		}
		return latestPrice.Price > alert.Price
	}
	return false
}
