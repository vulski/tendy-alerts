package notifiers

import (
	"fmt"
	tendy_alerts "github.com/vulski/tendy-alerts"
)

type EmailNotifier struct {
}

func NewEmailNotifier() *EmailNotifier {
	return &EmailNotifier{}
}

func (en *EmailNotifier) NotifyUser(currencyLog tendy_alerts.CurrencyPriceLog, alert tendy_alerts.Alert) error {
	fmt.Println("I sent an email!")
	return nil
}
