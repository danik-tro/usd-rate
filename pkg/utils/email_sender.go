package utils

import (
	"fmt"
	"log"

	"github.com/danik-tro/usd-rate/pkg/core"
	"gopkg.in/gomail.v2"
)

func SendMessage(config *core.Config, email string, rate float64) error {
	m := gomail.NewMessage()

	m.SetHeader("From", config.SMTPUsername)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Latest USD rate")
	m.SetBody("text/plain", fmt.Sprintf("Current rate: %f", rate))

	d := gomail.NewDialer(config.SMTPHost, 587, config.SMTPUsername, config.SMTPPassword)
	if err := d.DialAndSend(m); err != nil {
		log.Printf("Failed to send the email: %v\n", err)
		return err
	}

	return nil
}
