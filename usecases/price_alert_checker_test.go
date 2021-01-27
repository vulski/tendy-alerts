package usecases

import (
	tendy_alerts "github.com/vulski/tendy-alerts"
	"github.com/vulski/tendy-alerts/enums"
	"github.com/vulski/tendy-alerts/mocks"
	"testing"
	"github.com/golang/mock/gomock"
)

func TestItPullsUsers(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mocks.NewMockUserRepository(ctrl)

	// There is active users
	user := &tendy_alerts.User{Username: "test"}
	user.ID = 1
	userRepoMock.EXPECT().GetAllActive().Return([]*tendy_alerts.User{user}, nil)

	// And that user has active Alerts.
	targetAlert := tendy_alerts.Alert{
		Currency:         "BTC",
		Price:            10000,
		PercentageChange: 0,
		Type:             enums.AlertType_TARGET_ALERT,
		Frequency:        enums.AlertFrequency_ONE_TIME,
		Comparison:       enums.AlertComparison_GREATER_THAN,
		TradePair:        "BTC/USD",
		Active:           true,
		UserId:           user.ID,
	}
	alertRepoMock := mocks.NewMockAlertRepository(ctrl)
	alertRepoMock.EXPECT().GetAllForUserIDs([]uint{user.ID}).Return([]tendy_alerts.Alert{targetAlert}, nil)

	// Then the User should be notified.
	notifierMock := mocks.NewMockNotifier(ctrl)
	notifierMock.EXPECT().NotifyUser(tendy_alerts.CurrencyPriceLog{}, targetAlert).Return(nil)

	sut := NewPriceNotificationManager(userRepoMock, alertRepoMock, notifierMock)

	// When we check the User's Alerts
	sut.Run()
}
