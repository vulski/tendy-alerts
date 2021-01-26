package usecases

import (
	"github.com/vulski/tendy-alerts/enums"
	"github.com/vulski/tendy-alerts/models"
)

type AlertEvaluator struct {
	currencyLogRepository models.CurrencyPriceLogRepository
}

func NewAlertEvaluator(currencyLogRepo models.CurrencyPriceLogRepository) *AlertEvaluator {
	return &AlertEvaluator{currencyLogRepository: currencyLogRepo}
}

func (a *AlertEvaluator) ShouldAlertUser(latestPrice models.CurrencyPriceLog, alert *models.Alert) bool {
	if alert.Type == enums.AlertType_TARGET_ALERT {
		if alert.Comparison == enums.AlertComparison_LESS_THAN {
			return latestPrice.Price < alert.Price
		}
		return  latestPrice.Price > alert.Price
	}
	return false
}
