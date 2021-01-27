package tendy_alerts

//go:generate mockgen -destination=mocks/mock_notifier.go -package=mocks . Notifier
type Notifier interface {
	NotifyUser(currencyLog CurrencyPriceLog, alert Alert) error
}