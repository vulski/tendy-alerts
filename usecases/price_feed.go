package usecases

import "github.com/vulski/tendy-alerts/models"

var channels map[string]*chan models.CurrencyPriceLog

var quit chan interface{}

func init() {
	channels = make(map[string]*chan models.CurrencyPriceLog)
}

func SubscribeToCurrency(coin string) *chan models.CurrencyPriceLog {
	if coinChan, ok := channels[coin]; ok {
		return coinChan
	}
	coinChan := make(chan models.CurrencyPriceLog)
	channels[coin] = &coinChan
	return channels[coin]
}

func StartPriceFeed() {

}

func StopPriceFeed() {

}