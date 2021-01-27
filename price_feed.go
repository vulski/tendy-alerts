package tendy_alerts

type PriceFeed interface {
	GetCurrencyFeed(currency string) (*chan CurrencyPriceLog, error)
	StartFeed() error
	StopFeed() error
}
