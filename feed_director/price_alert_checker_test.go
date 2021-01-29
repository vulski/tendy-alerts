package feed_director

import (
	"github.com/golang/mock/gomock"
	ta "github.com/vulski/tendy-alerts"
	"github.com/vulski/tendy-alerts/mocks"
	"testing"
)

type priceCheckerMocks struct {
	notifierFactory *mocks.MockNotifierFactory
	alertRepo *mocks.MockAlertRepository
	priceRepo *mocks.MockPriceSnapshotRepository
}

func createPriceCheckerMocked(ctrl *gomock.Controller) (*PriceChecker, *priceCheckerMocks) {
	mks := &priceCheckerMocks{
		notifierFactory: mocks.NewMockNotifierFactory(ctrl),
		alertRepo:       mocks.NewMockAlertRepository(ctrl),
		priceRepo:       mocks.NewMockPriceSnapshotRepository(ctrl),
	}

	return NewPriceChecker(mks.notifierFactory, mks.alertRepo, mks.priceRepo), mks
}

func TestPriceChecker_CheckPrice_PercentageChangeAlertSuccess(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// And that there are active alerts.
	targetAlert := ta.Alert{
		Currency:             "BTC",
		Price:                0,
		PercentageChange:     .20,
		Type:                 ta.PercentageChangeAlert,
		Comparison:           ta.GreaterThanComparison,
		TradePair:            "BTC/USD",
		Active:               true,
		NotificationSettings: ta.NotificationSettings{Type: ta.EmailNotification},
	}
	targetAlert.ID = 3

	sut, mks := createPriceCheckerMocked(ctrl)

	latestPrice := ta.PriceSnapshot{Currency: "BTC", Price: 10001}
	alerts := []ta.Alert{targetAlert}
	mks.alertRepo.EXPECT().GetActiveAlertsForCurrency(latestPrice.Currency).Return(alerts, nil).Times(1)

	// Then the User should be notified.
	notifierMock := mocks.NewMockNotifier(ctrl)
	notifierMock.EXPECT().NotifyUser(latestPrice, targetAlert).Return(nil).Times(1)
	mks.notifierFactory.EXPECT().CreateNotifierFromType(ta.EmailNotification).Return(notifierMock, nil).Times(1)


	targetAlert.Active = false
	mks.alertRepo.EXPECT().Save(targetAlert).Return(&targetAlert, nil)

	// When we check the User's Alerts
	sut.CheckPrice(latestPrice)
}

func TestPriceChecker_CheckPrice_ItWillDisableOneTimeAlertsAfterSuccessfullyNotifying(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// And that there are active alerts.
	targetAlert := ta.Alert{
		Currency:             "BTC",
		Price:                10000,
		PercentageChange:     0,
		Type:                 ta.TargetAlert,
		Comparison:           ta.GreaterThanComparison,
		TradePair:            "BTC/USD",
		Active:               true,
		NotificationSettings: ta.NotificationSettings{Type: ta.EmailNotification},
	}
	targetAlert.ID = 3

	latestPrice := ta.PriceSnapshot{Currency: "BTC", Price: 10001}
	alerts := []ta.Alert{targetAlert}
	alertRepoMock := mocks.NewMockAlertRepository(ctrl)
	alertRepoMock.EXPECT().GetActiveAlertsForCurrency(latestPrice.Currency).Return(alerts, nil).Times(1)

	// Then the User should be notified.
	notifierMock := mocks.NewMockNotifier(ctrl)
	notifierMock.EXPECT().NotifyUser(latestPrice, targetAlert).Return(nil).Times(1)
	notifierFactoryMock := mocks.NewMockNotifierFactory(ctrl)
	notifierFactoryMock.EXPECT().CreateNotifierFromType(ta.EmailNotification).Return(notifierMock, nil).Times(1)

	priceRepoMock := mocks.NewMockPriceSnapshotRepository(ctrl)

	sut := NewPriceChecker(notifierFactoryMock, alertRepoMock, priceRepoMock)

	targetAlert.Active = false
	alertRepoMock.EXPECT().Save(targetAlert).Return(&targetAlert, nil)

	// When we check the User's Alerts
	sut.CheckPrice(latestPrice)
}

func TestPriceChecker_CheckPrice_WillGetsActiveAlertsForTheGivenCurrencyPriceLogAndWillNotifyTheUser(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// And that there are active alerts.
	targetAlert := ta.Alert{
		Currency:             "BTC",
		Price:                10000,
		PercentageChange:     0,
		Type:                 ta.TargetAlert,
		Comparison:           ta.GreaterThanComparison,
		TradePair:            "BTC/USD",
		Active:               true,
		NotificationSettings: ta.NotificationSettings{Type: ta.EmailNotification},
	}
	targetAlert.ID = 3

	latestPrice := ta.PriceSnapshot{Currency: "BTC", Price: 10001}
	alerts := []ta.Alert{targetAlert}
	alertRepoMock := mocks.NewMockAlertRepository(ctrl)
	alertRepoMock.EXPECT().GetActiveAlertsForCurrency(latestPrice.Currency).Return(alerts, nil).Times(1)

	// Then the User should be notified.
	notifierMock := mocks.NewMockNotifier(ctrl)
	notifierMock.EXPECT().NotifyUser(latestPrice, targetAlert).Return(nil).Times(1)
	notifierFactoryMock := mocks.NewMockNotifierFactory(ctrl)
	notifierFactoryMock.EXPECT().CreateNotifierFromType(ta.EmailNotification).Return(notifierMock, nil).Times(1)

	targetAlert.Active = false
	alertRepoMock.EXPECT().Save(targetAlert).Return(&targetAlert, nil)

	priceRepoMock := mocks.NewMockPriceSnapshotRepository(ctrl)

	sut := NewPriceChecker(notifierFactoryMock, alertRepoMock, priceRepoMock)

	// When we check the User's Alerts
	sut.CheckPrice(latestPrice)
}

func TestPriceChecker_CheckPrice_WillNotNotifyTheUserIfItShouldNot(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// And that there are active alerts.
	targetAlert := ta.Alert{
		Currency:             "BTC",
		Price:                10000,
		PercentageChange:     0,
		Type:                 ta.TargetAlert,
		Comparison:           ta.LessThanComparison,
		TradePair:            "BTC/USD",
		Active:               true,
		NotificationSettings: ta.NotificationSettings{Type: ta.EmailNotification},
	}
	targetAlert.ID = 3

	latestPrice := ta.PriceSnapshot{Currency: "BTC", Price: 10001}
	alerts := []ta.Alert{targetAlert}
	alertRepoMock := mocks.NewMockAlertRepository(ctrl)
	alertRepoMock.EXPECT().GetActiveAlertsForCurrency(latestPrice.Currency).Return(alerts, nil).Times(1)

	// Then the User should be notified.
	notifierFactoryMock := mocks.NewMockNotifierFactory(ctrl)
	notifierFactoryMock.EXPECT().CreateNotifierFromType(gomock.Any()).Times(0)

	alertRepoMock.EXPECT().Save(gomock.Any()).Times(0)

	priceRepoMock := mocks.NewMockPriceSnapshotRepository(ctrl)

	sut := NewPriceChecker(notifierFactoryMock, alertRepoMock, priceRepoMock)

	// When we check the User's Alerts
	sut.CheckPrice(latestPrice)
}

func TestPriceChecker_shouldAlertUser_PercentageChangeLessThanComparisonShouldPass(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	targetAlert := ta.Alert{
		Currency:         "BTC",
		PercentageChange: .20,
		Type:             ta.PercentageChangeAlert,
		Comparison:       ta.LessThanComparison,
		TradePair:        "BTC/USD",
		Active:           true,
	}
	latestPrice := ta.PriceSnapshot{Price: 20001}

	// When
	sut, _ := createPriceCheckerMocked(ctrl)

	// Then
	if true == sut.shouldAlertUser(latestPrice, targetAlert) {
		t.Fail()
	}
}

func TestPriceChecker_shouldAlertUser_WillIgnoreInActiveAlerts(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	targetAlert := ta.Alert{
		Currency:   "BTC",
		Price:      20000,
		Type:       ta.TargetAlert,
		Comparison: ta.LessThanComparison,
		TradePair:  "BTC/USD",
		Active:     false,
	}
	latestPrice := ta.PriceSnapshot{Price: 20001}

	// When
	sut, _ := createPriceCheckerMocked(ctrl)

	// Then
	if true == sut.shouldAlertUser(latestPrice, targetAlert) {
		t.Fail()
	}
}

func TestPriceChecker_shouldAlertUser_TargetAlertLessThanComparisonShouldPass(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	targetAlert := ta.Alert{
		Currency:   "BTC",
		Price:      20001,
		Type:       ta.TargetAlert,
		Comparison: ta.LessThanComparison,
		TradePair:  "BTC/USD",
		Active:     true,
	}
	latestPrice := ta.PriceSnapshot{Price: 20000}

	// When
	sut, _ := createPriceCheckerMocked(ctrl)

	// Then
	if false == sut.shouldAlertUser(latestPrice, targetAlert) {
		t.Fail()
	}
}

func TestPriceChecker_shouldAlertUser_TargetAlertGreaterThanComparisonShouldPass(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	targetAlert := ta.Alert{
		Currency:   "BTC",
		Price:      19999,
		Type:       ta.TargetAlert,
		Comparison: ta.GreaterThanComparison,
		TradePair:  "BTC/USD",
		Active:     true,
	}
	latestPrice := ta.PriceSnapshot{Price: 20000}

	// When
	sut, _ := createPriceCheckerMocked(ctrl)

	// Then
	if false == sut.shouldAlertUser(latestPrice, targetAlert) {
		t.Fail()
	}
}

func TestPriceChecker_shouldAlertUser_TargetAlertGreaterThanComparisonShouldFail(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	targetAlert := ta.Alert{
		Currency:   "BTC",
		Price:      20000,
		Type:       ta.TargetAlert,
		Comparison: ta.GreaterThanComparison,
		TradePair:  "BTC/USD",
		Active:     true,
	}
	latestPrice := ta.PriceSnapshot{Price: 19999}

	// When
	sut, _ := createPriceCheckerMocked(ctrl)

	// Then
	if true == sut.shouldAlertUser(latestPrice, targetAlert) {
		t.Fail()
	}
}
