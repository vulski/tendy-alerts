package notifiers

import tendy_alerts "github.com/vulski/tendy-alerts"

type EmailNotifier struct {
}

func NewEmailNotifier() *EmailNotifier {
	return &EmailNotifier{}
}

func (en *EmailNotifier) NotifyUser(currencyLog tendy_alerts.PriceSnapshot, alert tendy_alerts.Alert) error {
	panic("I'm an email!")
	return nil
}
