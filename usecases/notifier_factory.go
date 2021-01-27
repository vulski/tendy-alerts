package usecases

import (
	"errors"
	tendy "github.com/vulski/tendy-alerts"
	"github.com/vulski/tendy-alerts/enums"
	"github.com/vulski/tendy-alerts/usecases/notifiers"
)

type NotifierFactory struct {
	notifiers map[enums.NotificationType]tendy.Notifier
}

func NewNotifierFactory() *NotifierFactory {
	factory := &NotifierFactory{}
	factory.notifiers = make(map[enums.NotificationType]tendy.Notifier)
	factory.notifiers[enums.NotificationType_EMAIL] = notifiers.NewEmailNotifier()

	return factory
}

func (nf *NotifierFactory) CreateNotifierFromType(notificationType enums.NotificationType) (tendy.Notifier, error) {
	if notifier, ok := nf.notifiers[notificationType]; ok {
		return notifier, nil
	}

	return nil, errors.New("unknown NotificationType")
}
