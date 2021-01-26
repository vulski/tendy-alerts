package usecases

import (
	"github.com/vulski/tendy-alerts/models"
)

type AlertEvaluator struct {
	currencyLogRepository models.CurrencyPriceLogRepository
}

func NewAlertEvaluator(currencyLogRepo models.CurrencyPriceLogRepository) *AlertEvaluator {
	return &AlertEvaluator{currencyLogRepository: currencyLogRepo}
}

func (a *AlertEvaluator) ShouldAlertUser(latestPrice models.CurrencyPriceLog, alert *models.Alert) bool {
	if latestPrice.Price > alert.Price {
		return true
	}
	return false
}
