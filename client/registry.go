package client

import (
	"github.com/techx/portal/client/db"
	"github.com/techx/portal/client/twilio"
	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"gopkg.in/gomail.v2"
)

type Registry struct {
	DB     *db.Repository
	Twilio twilio.Client
	GMail  *gomail.Dialer
}

func NewRegistry(cfg config.Config) *Registry {
	usersDB, err := db.NewRepository(cfg, constants.TableNameUsers)
	if err != nil {
		panic(err)
	}

	twilioClient := twilio.NewTwilioClient(cfg.Twilio)
	gMailClient := gomail.NewDialer(
		cfg.GMail.SMTPServer,
		cfg.GMail.SMTPPort,
		cfg.GMail.From,
		cfg.GMail.SMTPPassword,
	)

	return &Registry{
		DB:     usersDB,
		Twilio: twilioClient,
		GMail:  gMailClient,
	}
}
