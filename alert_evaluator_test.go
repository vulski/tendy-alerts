package tendy_alerts

import (
	"testing"
)

var alertEvaluator *AlertEvaluator

func init() {
	alertEvaluator = NewAlertEvaluator()
}

type CurrencyPriceLogRepoStub struct {
	Logs []*CurrencyPriceLog
}

func (r *CurrencyPriceLogRepoStub) AddLog(log *CurrencyPriceLog) {
	r.Logs = append(r.Logs, log)
}

func (r *CurrencyPriceLogRepoStub) GetWithConstraints(constraints CurrencyPriceLogRepoConstraints) ([]*CurrencyPriceLog, error) {
	constraints.Name = ""
	return r.Logs, nil
}

func TestWillIgnoreInActiveAlerts(t *testing.T) {
	// Given
	targetAlert := Alert{
		Currency:   "BTC",
		Price:      20000,
		Type:       TargetAlert,
		Frequency:  OneTimeFrequency,
		Comparison: LessThanComparison,
		TradePair:  "BTC/USD",
		Active:     false,
	}
	latestPrice := CurrencyPriceLog{Price: 20001}

	// When
	// Then
	if true == alertEvaluator.ShouldAlertUser(latestPrice, targetAlert) {
		t.Fail()
	}
}

func TestTargetAlertLessThanComparisonShouldPass(t *testing.T) {
	// Given
	targetAlert := Alert{
		Currency:   "BTC",
		Price:      20001,
		Type:       TargetAlert,
		Frequency:  OneTimeFrequency,
		Comparison: LessThanComparison,
		TradePair:  "BTC/USD",
		Active:     true,
	}
	latestPrice := CurrencyPriceLog{Price: 20000}

	// When
	// Then
	if false == alertEvaluator.ShouldAlertUser(latestPrice, targetAlert) {
		t.Fail()
	}
}

func TestTargetAlertGreaterThanComparisonShouldPass(t *testing.T) {
	// Given
	targetAlert := Alert{
		Currency:   "BTC",
		Price:      19999,
		Type:       TargetAlert,
		Frequency:  OneTimeFrequency,
		Comparison: GreaterThanComparison,
		TradePair:  "BTC/USD",
		Active:     true,
	}
	latestPrice := CurrencyPriceLog{Price: 20000}

	// When
	// Then
	if false == alertEvaluator.ShouldAlertUser(latestPrice, targetAlert) {
		t.Fail()
	}
}

func TestTargetAlertGreaterThanComparisonShouldFail(t *testing.T) {
	// Given
	targetAlert := Alert{
		Currency:   "BTC",
		Price:      20000,
		Type:       TargetAlert,
		Frequency:  OneTimeFrequency,
		Comparison: GreaterThanComparison,
		TradePair:  "BTC/USD",
		Active:     true,
	}
	latestPrice := CurrencyPriceLog{Price: 19999}

	// When
	// Then
	if true == alertEvaluator.ShouldAlertUser(latestPrice, targetAlert) {
		t.Fail()
	}
}
