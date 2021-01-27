package usecases

import (
	tendy_alerts "github.com/vulski/tendy-alerts"
	"github.com/vulski/tendy-alerts/enums"
	"testing"
)

var userRepoMock tendy_alerts.UserRepoMock
var alertRepoMock tendy_alerts.AlertRepoMock

func TestItPullsUsers(t *testing.T) {
	// Given
	// There is active users
	userRepoMock.Users = []*tendy_alerts.User{{Username: "test"}}
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
	}
	alertRepoMock.Alerts = append(alertRepoMock.Alerts, targetAlert)

	sut := NewPriceNotificationManager(&userRepoMock)

	// When we check the User's Alerts
	sut.Run()

	// Then the User should be notified.

}
