package enums

type AlertFrequency struct{ value string }

func (af *AlertFrequency) Value() string {
	return af.value
}

func (af *AlertFrequency) Is(frequency AlertFrequency) bool {
	return frequency.Value() == af.Value()
}

func AlertFrequencyOneTime() AlertFrequency {
	return AlertFrequency{value: "one_time"}
}

func AlertFrequency15m() AlertFrequency {
	return AlertFrequency{value: "15m"}
}

func AlertFrequency30m() AlertFrequency {
	return AlertFrequency{value: "30m"}
}

func AlertFrequency1h() AlertFrequency {
	return AlertFrequency{value: "1h"}
}

func AlertFrequency6hr() AlertFrequency {
	return AlertFrequency{value: "6h"}
}

func AlertFrequency12hr() AlertFrequency {
	return AlertFrequency{value: "12h"}
}

func AlertFrequency24hr() AlertFrequency {
	return AlertFrequency{value: "24h"}
}
