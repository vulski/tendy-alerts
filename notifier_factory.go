package tendy_alerts

//go:generate mockgen -destination=mocks/mock_notifier_factory.go -package=mocks . NotifierFactory
type NotifierFactory interface {
	CreateNotifierFromType(notificationType NotificationType) (Notifier, error)
}
