package services

import (
	"github.com/vulski/tendy-alerts/enums"
	"github.com/vulski/tendy-alerts/models"
	"testing"
)

var alertEvaluator *AlertEvaluator

func init() {
	alertEvaluator = NewAlertEvaluator()
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

func TestWillIgnoreInActiveAlerts(t *testing.T) {
	// Given
	targetAlert := models.Alert{
		Currency:   "BTC",
		Price:      20000,
		Type:       enums.AlertType_TARGET_ALERT,
		Frequency:  enums.AlertFrequency_ONE_TIME,
		Comparison: enums.AlertComparison_LESS_THAN,
		TradePair:  "BTC/USD",
		Active:     false,
	}
	latestPrice := models.CurrencyPriceLog{Price: 20001}

	// When
	// Then
	if true == alertEvaluator.ShouldAlertUser(latestPrice, targetAlert) {
		t.Fail()
	}
}

func TestTargetAlertLessThanComparisonShouldPass(t *testing.T) {
	// Given
	targetAlert := models.Alert{
		Currency:   "BTC",
		Price:      20000,
		Type:       enums.AlertType_TARGET_ALERT,
		Frequency:  enums.AlertFrequency_ONE_TIME,
		Comparison: enums.AlertComparison_LESS_THAN,
		TradePair:  "BTC/USD",
		Active:     true,
	}
	latestPrice := models.CurrencyPriceLog{Price: 20001}

	// When
	// Then
	if false == alertEvaluator.ShouldAlertUser(latestPrice, targetAlert) {
		t.Fail()
	}
}

func TestTargetAlertGreaterThanComparisonShouldPass(t *testing.T) {
	// Given
	targetAlert := models.Alert{
		Currency:   "BTC",
		Price:      20000,
		Type:       enums.AlertType_TARGET_ALERT,
		Frequency:  enums.AlertFrequency_ONE_TIME,
		Comparison: enums.AlertComparison_GREATER_THAN,
		TradePair:  "BTC/USD",
		Active:     true,
	}
	latestPrice := models.CurrencyPriceLog{Price: 19999}

	// When
	// Then
	if false == alertEvaluator.ShouldAlertUser(latestPrice, targetAlert) {
		t.Fail()
	}
}

func TestTargetAlertGreaterThanComparisonShouldFail(t *testing.T) {
	// Given
	targetAlert := models.Alert{
		Currency:   "BTC",
		Price:      20000,
		Type:       enums.AlertType_TARGET_ALERT,
		Frequency:  enums.AlertFrequency_ONE_TIME,
		Comparison: enums.AlertComparison_GREATER_THAN,
		TradePair:  "BTC/USD",
		Active:     true,
	}
	latestPrice := models.CurrencyPriceLog{Price: 20001}

	// When
	// Then
	if true == alertEvaluator.ShouldAlertUser(latestPrice, targetAlert) {
		t.Fail()
	}
}
