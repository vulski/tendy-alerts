package feed_director

import (
	ta "github.com/vulski/tendy-alerts"
	"time"
)

// TODO: Better name lel
type Director struct {
	running      bool
	feeds        map[string]map[string]chan ta.PriceSnapshot
	exchanges    []ta.PriceFeed
	priceChecker PriceChecker
}

func New(priceChecker PriceChecker, exchanges []ta.PriceFeed) Director {
	return Director{
		exchanges:    exchanges,
		feeds:        make(map[string]map[string]chan ta.PriceSnapshot),
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
			m.feeds[exchange.ExchangeName()] = make(map[string]chan ta.PriceSnapshot)
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

func (m *Director) processFeed(feed chan ta.PriceSnapshot) {
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
