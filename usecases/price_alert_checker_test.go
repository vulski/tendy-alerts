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
	targetAlert.ID = 3
	alertRepoMock := mocks.NewMockAlertRepository(ctrl)
	alertRepoMock.EXPECT().GetAllForUserIDs([]uint{user.ID}).Return([]tendy.Alert{targetAlert}, nil)

	// Then the User should be notified.
	notificationSettingsRepoMock := mocks.NewMockNotificationSettingRepository(ctrl)
	notificationSettingsRepoMock.EXPECT().GetForAlertId(targetAlert.ID).Return(
		tendy.NotificationSetting{Type: enums.NotificationType_EMAIL},
		nil,
	)

	notifierMock := mocks.NewMockNotifier(ctrl)
	notifierMock.EXPECT().NotifyUser(tendy.CurrencyPriceLog{}, targetAlert).Return(nil)
	notifierFactoryMock := mocks.NewMockNotifierFactory(ctrl)
	notifierFactoryMock.EXPECT().CreateNotifierFromType(enums.NotificationType_EMAIL).Return(notifierMock, nil)

	sut := NewPriceNotificationManager(notifierFactoryMock, userRepoMock, alertRepoMock, notificationSettingsRepoMock)

	// When we check the User's Alerts
	sut.Run()
}
