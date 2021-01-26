package usecases

import (
	"github.com/vulski/tendy-alerts/models"
)

var currencyLogRepository models.CurrencyPriceLogRepository

func ShouldAlertUser(alert *models.Alert) bool {

	return true
}