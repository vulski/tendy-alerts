package feed_director

import (
	tendy_alerts "github.com/vulski/tendy-alerts"
	"github.com/vulski/tendy-alerts/price_feeds"
	"time"
)

// TODO: Better name lel
type Director struct {
	running      bool
	feeds        map[string]map[string]chan tendy_alerts.PriceSnapshot
	exchanges    []tendy_alerts.PriceFeed
	priceChecker *PriceChecker
}

func New(priceChecker *PriceChecker) Director {
	return Director{
		exchanges:    []tendy_alerts.PriceFeed{price_feeds.NewCoinBasePriceFeed()},
		feeds:        make(map[string]map[string]chan tendy_alerts.PriceSnapshot),
		priceChecker: priceChecker,
	}
}

func (m *Director) Start() {
	m.running = true
	// Initialize exchange subscriptions.
	for _, exchange := range m.exchanges {
		// TODO: Add list of currencies? User based + Exchange based?
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

func (m *Director) Stop() {
	m.running = false
	for _, exchange := range m.exchanges {
		exchange.StopFeed()
	}
}

func (m *Director) processFeed(feed chan tendy_alerts.PriceSnapshot) {
	start := time.Now()
	for m.running {
		now := time.Now()
		select {
		case snapshot := <-feed:
			// TODO: Add a different way to rate limit.
			if now.Sub(start).Seconds() > 2 {
				start = time.Now()
				m.priceChecker.CheckPrice(snapshot)
			}
		}
	}
}
