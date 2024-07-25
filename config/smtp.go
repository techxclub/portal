package config

import (
	"fmt"

	"github.com/techx/portal/utils"
)

type MailSMTP struct {
	Retries      int    `yaml:"RETRIES" env:"RETRIES"`
	SMTPServer   string `yaml:"SMTP_SERVER" env:"SMTP_SERVER"`
	SMTPPort     int    `yaml:"SMTP_PORT" env:"SMTP_PORT"`
	SMTPUsername string `yaml:"SMTP_USERNAME" env:"SMTP_USERNAME"`
	SMTPPassword string `yaml:"SMTP_PASSWORD" env:"SMTP_PASSWORD"`
	Domain       string `yaml:"DOMAIN" env:"DOMAIN"`
	SenderEmail  string `yaml:"SENDER_EMAIL" env:"SENDER_EMAIL"`
}

func (m MailSMTP) GetSender(senderName string) string {
	return fmt.Sprintf("%s <%s>", senderName, m.SenderEmail)
}

func (m MailSMTP) GetMessageID() string {
	return fmt.Sprintf("<%s@%s>", utils.GetRandomUUID(), m.Domain)
}
