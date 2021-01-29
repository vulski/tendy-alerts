package notifiers

import (
	"encoding/json"
	"errors"
	"github.com/vulski/tendy-alerts"
	"io/ioutil"
)

type Factory struct {
	notifiers map[tendy_alerts.NotificationType]tendy_alerts.Notifier
}

func NewNotifierFactoryFromConfig(configPath string) (*Factory, error) {
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return &Factory{}, err
	}
	conf := Config{}
	err = json.Unmarshal(file, &conf)
	if err != nil {
		return &Factory{}, err
	}
	return NewNotifierFactory(conf), nil
}

func NewNotifierFactory(config Config) *Factory {
	factory := &Factory{}
	factory.notifiers = make(map[tendy_alerts.NotificationType]tendy_alerts.Notifier)
	factory.notifiers[tendy_alerts.EmailNotification] = &EmailNotifier{config: config.Email}

	return factory
}

func (nf *Factory) CreateNotifierFromType(notificationType tendy_alerts.NotificationType) (tendy_alerts.Notifier, error) {
	if notifier, ok := nf.notifiers[notificationType]; ok {
		return notifier, nil
	}

	return nil, errors.New("unknown NotificationType")
}
