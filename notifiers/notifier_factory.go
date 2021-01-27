package notifiers

import (
	"errors"
	"github.com/vulski/tendy-alerts"
)

type Factory struct {
	notifiers map[tendy_alerts.NotificationType]tendy_alerts.Notifier
}

func NewNotifierFactory() *Factory {
	factory := &Factory{}
	factory.notifiers = make(map[tendy_alerts.NotificationType]tendy_alerts.Notifier)
	factory.notifiers[tendy_alerts.EmailNotification] = NewEmailNotifier()

	return factory
}

func (nf *Factory) CreateNotifierFromType(notificationType tendy_alerts.NotificationType) (tendy_alerts.Notifier, error) {
	if notifier, ok := nf.notifiers[notificationType]; ok {
		return notifier, nil
	}

	return nil, errors.New("unknown NotificationType")
}
