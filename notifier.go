package tendy_alerts

type Notifier interface {
	NotifyUser(currencyLog CurrencyPriceLog, alert Alert, settings NotificationSetting) error
}