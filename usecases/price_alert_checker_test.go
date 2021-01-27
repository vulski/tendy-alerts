package usecases

import (
	"github.com/golang/mock/gomock"
	tendy "github.com/vulski/tendy-alerts"
	"github.com/vulski/tendy-alerts/enums"
	"github.com/vulski/tendy-alerts/mocks"
	"testing"
)

func TestItPullsUsers(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mocks.NewMockUserRepository(ctrl)

	// There is active users
	user := &tendy.User{Username: "test"}
	user.ID = 1
	userRepoMock.EXPECT().GetAllActive().Return([]*tendy.User{user}, nil)

	// And that user has active Alerts.
	targetAlert := tendy.Alert{
		Currency:             "BTC",
		Price:                10000,
		PercentageChange:     0,
		Type:                 enums.AlertType_TARGET_ALERT,
		Frequency:            enums.AlertFrequency_ONE_TIME,
		Comparison:           enums.AlertComparison_GREATER_THAN,
		TradePair:            "BTC/USD",
		Active:               true,
		UserId:               user.ID,
		NotificationSettings: tendy.NotificationSetting{Type: enums.NotificationType_EMAIL},
	}
	targetAlert.ID = 3
	user.Alerts = append(user.Alerts, targetAlert)

	// Then the User should be notified.
	notifierMock := mocks.NewMockNotifier(ctrl)
	notifierMock.EXPECT().NotifyUser(tendy.CurrencyPriceLog{}, targetAlert).Return(nil)
	notifierFactoryMock := mocks.NewMockNotifierFactory(ctrl)
	notifierFactoryMock.EXPECT().CreateNotifierFromType(enums.NotificationType_EMAIL).Return(notifierMock, nil)

	sut := NewPriceNotificationManager(notifierFactoryMock, userRepoMock)

	// When we check the User's Alerts
	sut.Run()
}
