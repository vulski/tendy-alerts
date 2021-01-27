package tendy_alerts

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
