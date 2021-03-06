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
	addr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	err := smtp.SendMail(addr, auth, conf.FromAddress, to, msg)
	if err != nil {
		return err
	}
	fmt.Println("Finished with no errors.")
	return nil
}
