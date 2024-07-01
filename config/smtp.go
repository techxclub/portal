package config

import (
	"fmt"

	"github.com/techx/portal/utils"
)

type MailSMTP struct {
	SMTPServer   string `yaml:"SMTP_SERVER" env:"SMTP_SERVER"`
	SMTPPort     int    `yaml:"SMTP_PORT" env:"SMTP_PORT"`
	SMTPUsername string `yaml:"SMTP_USERNAME" env:"SMTP_USERNAME"`
	SMTPPassword string `yaml:"SMTP_PASSWORD" env:"SMTP_PASSWORD"`
	Domain       string `yaml:"DOMAIN" env:"DOMAIN"`
	FromName     string `yaml:"FROM_NAME" env:"FROM_NAME"`
	FromEmail    string `yaml:"FROM_EMAIL" env:"FROM_EMAIL"`
}

func (m MailSMTP) GetFrom() string {
	return fmt.Sprintf("%s <%s>", m.FromName, m.FromEmail)
}

func (m MailSMTP) GetMessageID() string {
	return fmt.Sprintf("<%s@%s>", utils.GetRandomUUID(), m.Domain)
}
