package manager

import (
	"github.com/golang/mock/gomock"
	tendy "github.com/vulski/tendy-alerts"
	"github.com/vulski/tendy-alerts/mocks"
	"testing"
)

func TestItGetsActiveAlertsForTheGivenCurrencyPriceLogAndWillNotifyTheUser(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// And that there are active alerts.
	targetAlert := tendy.Alert{
		Currency:             "BTC",
		Price:                10000,
		PercentageChange:     0,
		Type:                 tendy.TargetAlert,
		Frequency:            tendy.OneTimeFrequency,
		Comparison:           tendy.GreaterThanComparison,
		TradePair:            "BTC/USD",
		Active:               true,
		NotificationSettings: tendy.NotificationSetting{Type: tendy.EmailNotification},
	}
	targetAlert.ID = 3

	latestPrice := tendy.PriceSnapshot{Currency: "BTC", Price: 10001}
	alerts := []tendy.Alert{targetAlert}
	alertRepoMock := mocks.NewMockAlertRepository(ctrl)
	alertRepoMock.EXPECT().GetActiveAlertsForCurrency(latestPrice.Currency).Return(alerts, nil).Times(1)

	// Then the User should be notified.
	notifierMock := mocks.NewMockNotifier(ctrl)
	notifierMock.EXPECT().NotifyUser(latestPrice, targetAlert).Return(nil).Times(1)
	notifierFactoryMock := mocks.NewMockNotifierFactory(ctrl)
	notifierFactoryMock.EXPECT().CreateNotifierFromType(tendy.EmailNotification).Return(notifierMock, nil).Times(1)

	sut := NewPriceAlertChecker(notifierFactoryMock, alertRepoMock)

	// When we check the User's Alerts
	sut.CheckPrice(latestPrice)
}


func TestItWillNotNotifyTheUserIfItShouldNot(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// And that there are active alerts.
	targetAlert := tendy.Alert{
		Currency:             "BTC",
		Price:                10000,
		PercentageChange:     0,
		Type:                 tendy.TargetAlert,
		Frequency:            tendy.OneTimeFrequency,
		Comparison:           tendy.LessThanComparison,
		TradePair:            "BTC/USD",
		Active:               true,
		NotificationSettings: tendy.NotificationSetting{Type: tendy.EmailNotification},
	}
	targetAlert.ID = 3

	latestPrice := tendy.PriceSnapshot{Currency: "BTC", Price: 10001}
	alerts := []tendy.Alert{targetAlert}
	alertRepoMock := mocks.NewMockAlertRepository(ctrl)
	alertRepoMock.EXPECT().GetActiveAlertsForCurrency(latestPrice.Currency).Return(alerts, nil).Times(1)

	// Then the User should be notified.
	notifierFactoryMock := mocks.NewMockNotifierFactory(ctrl)
	notifierFactoryMock.EXPECT().CreateNotifierFromType(gomock.Any()).Times(0)

	sut := NewPriceAlertChecker(notifierFactoryMock, alertRepoMock)

	// When we check the User's Alerts
	sut.CheckPrice(latestPrice)
}
