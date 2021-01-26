package usecases

import (
	"github.com/vulski/tendy-alerts/models"
	"testing"
)

var currencyPriceLogRepoStub CurrencyPriceLogRepoStub

func init() {
	currencyLogRepository = &currencyPriceLogRepoStub
}

type CurrencyPriceLogRepoStub struct {
	Logs []*models.CurrencyPriceLog
}

func (r *CurrencyPriceLogRepoStub) AddLog(log *models.CurrencyPriceLog) {
	r.Logs = append(r.Logs, log)
}

func (r *CurrencyPriceLogRepoStub) GetWithConstraints(constraints models.CurrencyPriceLogRepoConstraints) ([]*models.CurrencyPriceLog, error) {
	constraints.Name = ""
	return r.Logs, nil
}

func TestTargetAlert(t *testing.T) {
	// Given
	targetAlert := &models.Alert{
		Frequency: models.OneTime,
		Price: "20000",
		TradePair: "BTC/USD",
		Type: models.TargetAlert,
	}
	currencyPriceLogRepoStub.AddLog(&models.CurrencyPriceLog{Price: "2000"})

	// When
	// Then
	if false == ShouldAlertUser(targetAlert) {
		t.Fail()
	}
}
