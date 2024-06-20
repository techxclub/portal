package client

import (
	"github.com/techx/portal/client/cache"
	"github.com/techx/portal/client/db"
	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"gopkg.in/gomail.v2"
)

type Registry struct {
	DB       *db.Repository
	GMail    *gomail.Dialer
	OTPCache cache.Cache[string]
}

func NewRegistry(cfg *config.Config) *Registry {
	usersDB, err := db.NewRepository(cfg, constants.TableNameUsers)
	if err != nil {
		panic(err)
	}

	gMailClient := gomail.NewDialer(
		cfg.GMail.SMTPServer,
		cfg.GMail.SMTPPort,
		cfg.GMail.From,
		cfg.GMail.SMTPPassword,
	)
	otpCache := cache.NewCache()
	return &Registry{
		DB:       usersDB,
		GMail:    gMailClient,
		OTPCache: otpCache,
	}
}
