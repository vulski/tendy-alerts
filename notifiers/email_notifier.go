package notifiers

import (
	"fmt"
	tendy_alerts "github.com/vulski/tendy-alerts"
	"log"
	"net/smtp"
)

type EmailNotifier struct {
	config EmailConfiguration
}

type EmailConfiguration struct {
	Host        string `json:"host"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Port        uint   `json:"port"`
	FromAddress string `json:"fromAddress"`
	Debug       bool   `json:"debug"`
}

func NewEmailNotifier(config EmailConfiguration) *EmailNotifier {
	return &EmailNotifier{config: config}
}

func (en *EmailNotifier) NotifyUser(price tendy_alerts.PriceSnapshot, alert tendy_alerts.Alert) error {
	conf := en.config
	log.Println("Sending Email.")
	if conf.Debug {
		line := fmt.Sprintf("Email: To: %s From: %s Msg: %s", alert.NotificationSettings.TargetUsername, en.config.FromAddress, price.Stringify())
		log.Println(line)
		fmt.Println(line)
		return nil
	}
	auth := smtp.PlainAuth("", conf.Username, conf.Password, conf.Host)
	to := []string{alert.NotificationSettings.TargetUsername}
	msg := []byte(price.Stringify())
	err := smtp.SendMail(fmt.Sprintf("%s:%i", conf.Host, conf.Port), auth, conf.FromAddress, to, msg)
	if err != nil {
		return err
	}
	return nil
}
