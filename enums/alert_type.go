package enums

type AlertType struct{ value string }

func AlertTypeTargetAlert() AlertType {
	return AlertType{value: "target_alert"}
}
func AlertTypePercentageChangeAlert() AlertType {
	return AlertType{value: "percentage_change_alert"}
}
func (a *AlertType) Value() string {
	return a.value
}

func (a *AlertType) Is(alertType AlertType) bool {
	return alertType.Value() == a.Value()
}
