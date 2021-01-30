package feed_director

import (
	"github.com/golang/mock/gomock"
	ta "github.com/vulski/tendy-alerts"
	"github.com/vulski/tendy-alerts/mocks"
	"testing"
	"time"
)

type PriceCheckerStub struct {
	pricesChecked []ta.PriceSnapshot
}

func (p *PriceCheckerStub) CheckPrice(price ta.PriceSnapshot) error {
	p.pricesChecked = append(p.pricesChecked, price)
	return nil
}

func TestDirector_Start_WillSubscribeAllExchangesToBTCAndCreateABTCFeed(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	priceChecker, _ := createPriceCheckerMocked(ctrl)
	exchange := mocks.NewMockPriceFeed(ctrl)
	exchange.EXPECT().ExchangeName().Return("mock_exchange").Times(3)
	exchange.EXPECT().SubscribeToCurrency("BTC").Return(make(chan ta.PriceSnapshot), nil).Times(1)
	exchange.EXPECT().StartFeed().Times(1)
	sut := New(priceChecker, []ta.PriceFeed{exchange})

	// When
	sut.Start()

	// Then
	if len(sut.feeds) <= 0 {
		t.Fail()
	}

	exchangeFeed, ok := sut.feeds["mock_exchange"]
	if !ok || exchangeFeed == nil {
		t.Error("Exchange feed does not exist")
	}
	btcFeed, ok := exchangeFeed["BTC"]
	if !ok || btcFeed == nil {
		t.Error("BTC Feed does not exit.")
	}
	if sut.running == false {
		t.Error("running is not true")
	}
}

func TestDirector_Start_WillRunAGoRoutineThatChecksPricesEveryTwoSeconds(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	exchange := mocks.NewMockPriceFeed(ctrl)
	priceFeed := make(chan ta.PriceSnapshot)
	exchange.EXPECT().ExchangeName().Return("mock_exchange").Times(3)
	exchange.EXPECT().SubscribeToCurrency("BTC").Return(priceFeed, nil).Times(1)
	exchange.EXPECT().StartFeed().Times(1)
	priceChecker := &PriceCheckerStub{}
	sut := New(priceChecker, []ta.PriceFeed{exchange})

	// When
	sut.Start()

	priceFeed <-ta.PriceSnapshot{Price: 10_000}

	// Then
	time.Sleep(time.Second * 3)
	if len(priceChecker.pricesChecked) <= 0 {
		t.Error("no prices were checked.")
	}
}

