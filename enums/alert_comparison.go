package enums

type AlertComparison struct {
	value string
}

func (ac *AlertComparison) Value() string {
	return ac.value
}

func (ac *AlertComparison) Is(comparison AlertComparison) bool {
	return comparison.Value() == ac.Value()
}

func AlertComparisonGreaterThan() AlertComparison {
	return AlertComparison{value: "greater_than"}
}

func AlertComparisonLessThan() AlertComparison {
	return AlertComparison{value: "less_than"}
}
