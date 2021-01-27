package usecases

import (
	"github.com/golang/mock/gomock"
	tendy "github.com/vulski/tendy-alerts"
	"github.com/vulski/tendy-alerts/enums"
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
		Type:                 enums.AlertType_TARGET_ALERT,
		Frequency:            enums.AlertFrequency_ONE_TIME,
		Comparison:           enums.AlertComparison_GREATER_THAN,
		TradePair:            "BTC/USD",
		Active:               true,
		NotificationSettings: tendy.NotificationSetting{Type: enums.NotificationType_EMAIL},
	}
	targetAlert.ID = 3

	latestPrice := tendy.CurrencyPriceLog{Currency: "BTC", Price: 10001}
	alerts := []tendy.Alert{targetAlert}
	alertRepoMock := mocks.NewMockAlertRepository(ctrl)
	alertRepoMock.EXPECT().GetActiveAlertsForCurrency(latestPrice.Currency).Return(alerts, nil).Times(1)

	// Then the User should be notified.
	notifierMock := mocks.NewMockNotifier(ctrl)
	notifierMock.EXPECT().NotifyUser(latestPrice, targetAlert).Return(nil).Times(1)
	notifierFactoryMock := mocks.NewMockNotifierFactory(ctrl)
	notifierFactoryMock.EXPECT().CreateNotifierFromType(enums.NotificationType_EMAIL).Return(notifierMock, nil).Times(1)

	sut := NewPriceNotificationManager(notifierFactoryMock, alertRepoMock)

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
		Type:                 enums.AlertType_TARGET_ALERT,
		Frequency:            enums.AlertFrequency_ONE_TIME,
		Comparison:           enums.AlertComparison_LESS_THAN,
		TradePair:            "BTC/USD",
		Active:               true,
		NotificationSettings: tendy.NotificationSetting{Type: enums.NotificationType_EMAIL},
	}
	targetAlert.ID = 3

	latestPrice := tendy.CurrencyPriceLog{Currency: "BTC", Price: 10001}
	alerts := []tendy.Alert{targetAlert}
	alertRepoMock := mocks.NewMockAlertRepository(ctrl)
	alertRepoMock.EXPECT().GetActiveAlertsForCurrency(latestPrice.Currency).Return(alerts, nil).Times(1)

	// Then the User should be notified.
	notifierFactoryMock := mocks.NewMockNotifierFactory(ctrl)
	notifierFactoryMock.EXPECT().CreateNotifierFromType(gomock.Any()).Times(0)

	sut := NewPriceNotificationManager(notifierFactoryMock, alertRepoMock)

	// When we check the User's Alerts
	sut.CheckPrice(latestPrice)
}
