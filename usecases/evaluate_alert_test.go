package usecases

import (
	"github.com/vulski/tendy-alerts/models"
	"testing"
	"time"
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

func TestTargetAlertShouldNotNotify(t *testing.T) {
	// Given
	targetAlert := &models.Alert{
		Currency:  "BTC",
		Frequency: models.OneTime,
		Price:     "20000",
		TradePair: "BTC/USD",
		Type:      models.TargetAlert,
	}
	thirtyMinutesAgo := time.Now().Add(time.Duration(-30) * time.Minute)
	currencyPriceLogRepoStub.AddLog(&models.CurrencyPriceLog{
		Currency:  "BTC",
		Price:     "10000",
		Timestamp: thirtyMinutesAgo,
	})

	// When
	// Then
	if true == ShouldAlertUser(targetAlert) {
		t.Fail()
	}
}

func TestTargetAlertSuccess(t *testing.T) {
	// Given
	targetAlert := &models.Alert{
		Currency:  "BTC",
		Frequency: models.OneTime,
		Price:     "20000",
		TradePair: "BTC/USD",
		Type:      models.TargetAlert,
	}
	thirtyMinutesAgo := time.Now().Add(time.Duration(-30) * time.Minute)
	currencyPriceLogRepoStub.AddLog(&models.CurrencyPriceLog{
		Currency:  "BTC",
		Price:     "20000",
		Timestamp: thirtyMinutesAgo,
	})

	// When
	// Then
	if false == ShouldAlertUser(targetAlert) {
		t.Fail()
	}
}
