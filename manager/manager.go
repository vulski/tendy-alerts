package manager

import (
	"fmt"
	tendy_alerts "github.com/vulski/tendy-alerts"
	"github.com/vulski/tendy-alerts/price_feeds"
)

type Manager struct {
	running      bool
	feeds        map[string]map[string]chan tendy_alerts.PriceSnapshot
	exchanges    []tendy_alerts.PriceFeed
	priceChecker *PriceChecker
}

func New(priceChecker *PriceChecker) Manager {
	return Manager{
		exchanges:    []tendy_alerts.PriceFeed{price_feeds.NewCoinBasePriceFeed()},
		feeds:        make(map[string]map[string]chan tendy_alerts.PriceSnapshot),
		priceChecker: priceChecker,
	}
}

func (m *Manager) Start() {
	m.running = true
	// Initialize exchange subscriptions.
	for _, exchange := range m.exchanges {
		// Default to BTC for now.
		currency := "BTC"
		feed, err := exchange.SubscribeToCurrency(currency)
		if err != nil {
			panic(err)
		}
		if m.feeds[exchange.ExchangeName()] == nil {
			m.feeds[exchange.ExchangeName()] = make(map[string]chan tendy_alerts.PriceSnapshot)
		}
		m.feeds[exchange.ExchangeName()][currency] = feed
		go m.processFeed(feed)
		exchange.StartFeed()
	}
}

func (m *Manager) Stop() {
	m.running = false
	for _, exchange := range m.exchanges {
		exchange.StopFeed()
	}
}

func (m *Manager) processFeed(feed chan tendy_alerts.PriceSnapshot) {
	for m.running {
		fmt.Println("Checking coinbase")
		select {
		case snapshot := <-feed:
			fmt.Println("New price to check!")
			m.priceChecker.CheckPrice(snapshot)
		}
	}
}
