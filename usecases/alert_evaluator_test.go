package usecases

import (
	"github.com/vulski/tendy-alerts/enums"
	"github.com/vulski/tendy-alerts/models"
	"testing"
)

var sut *AlertEvaluator
var currencyPriceLogRepoStub CurrencyPriceLogRepoStub

func init() {
	sut = NewAlertEvaluator(&currencyPriceLogRepoStub)
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

func TestTargetAlertLessThanComparisonShouldPass(t *testing.T) {
	// Given
	targetAlert := &models.Alert{
		Currency:  "BTC",
		Frequency: enums.AlertFrequencyOneTime(),
		Price:     20000,
		TradePair: "BTC/USD",
		Type:      enums.AlertTypeTargetAlert(),
	}
	latestPrice := models.CurrencyPriceLog{Price: 20001}

	// When
	// Then
	if false == sut.ShouldAlertUser(latestPrice, targetAlert) {
		t.Fail()
	}
}

func TestTargetAlertGreaterThanComparisonShouldPass(t *testing.T) {
	// Given
	targetAlert := &models.Alert{
		Currency:  "BTC",
		Frequency: enums.AlertFrequencyOneTime(),
		Price:     20000,
		TradePair: "BTC/USD",
		Type:      enums.AlertTypeTargetAlert(),
	}
	latestPrice := models.CurrencyPriceLog{Price: 20001}

	// When
	// Then
	if false == sut.ShouldAlertUser(latestPrice, targetAlert) {
		t.Fail()
	}
}

func TestTargetAlertGreaterThanComparisonShouldFail(t *testing.T) {
	// Given
	targetAlert := &models.Alert{
		Currency:  "BTC",
		Frequency: enums.AlertFrequencyOneTime(),
		Price:     20000,
		TradePair: "BTC/USD",
		Type:      enums.AlertTypeTargetAlert(),
		Comparison: enums.AlertComparisonGreaterThan(),
	}
	latestPrice := models.CurrencyPriceLog{Price: 10000}

	// When
	// Then
	if true == sut.ShouldAlertUser(latestPrice, targetAlert) {
		t.Fail()
	}
}