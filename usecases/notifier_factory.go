package usecases

import (
	"errors"
	tendy "github.com/vulski/tendy-alerts"
	"github.com/vulski/tendy-alerts/usecases/notifiers"
)

type NotifierFactory struct {
	notifiers map[tendy.NotificationType]tendy.Notifier
}

func NewNotifierFactory() *NotifierFactory {
	factory := &NotifierFactory{}
	factory.notifiers = make(map[tendy.NotificationType]tendy.Notifier)
	factory.notifiers[tendy.EmailNotification] = notifiers.NewEmailNotifier()

	return factory
}

func (nf *NotifierFactory) CreateNotifierFromType(notificationType tendy.NotificationType) (tendy.Notifier, error) {
	if notifier, ok := nf.notifiers[notificationType]; ok {
		return notifier, nil
	}

	return nil, errors.New("unknown NotificationType")
}
