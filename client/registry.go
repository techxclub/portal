package client

import (
	"github.com/techx/portal/client/db"
	"github.com/techx/portal/client/twilio"
	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
)

type Registry struct {
	DB     *db.Repository
	Twilio twilio.Client
}

func NewRegistry(cfg config.Config) *Registry {
	usersDB, err := db.NewRepository(cfg, constants.TableNameUsers)
	if err != nil {
		panic(err)
	}

	twilioClient := twilio.NewTwilioClient(cfg.Twilio)

	return &Registry{
		DB:     usersDB,
		Twilio: twilioClient,
	}
}
