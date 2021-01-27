package tendy_alerts

import (
	"testing"
)

var alertEvaluator *AlertEvaluator

func init() {
	alertEvaluator = NewAlertEvaluator()
}

type CurrencyPriceLogRepoStub struct {
	Logs []*PriceSnapshot
}

func (r *CurrencyPriceLogRepoStub) AddLog(log *PriceSnapshot) {
	r.Logs = append(r.Logs, log)
}

func (r *CurrencyPriceLogRepoStub) GetWithConstraints(constraints CurrencyPriceLogRepoConstraints) ([]*PriceSnapshot, error) {
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
	latestPrice := PriceSnapshot{Price: 20001}

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
	latestPrice := PriceSnapshot{Price: 20000}

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
	latestPrice := PriceSnapshot{Price: 20000}

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
	latestPrice := PriceSnapshot{Price: 19999}

	// When
	// Then
	if true == alertEvaluator.ShouldAlertUser(latestPrice, targetAlert) {
		t.Fail()
	}
}
