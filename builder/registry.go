package builder

import (
	"github.com/techx/portal/client"
	"github.com/techx/portal/config"
)

type Registry struct {
	UserInfoBuilder UserInfoBuilder
	UserAuthBuilder UserAuthBuilder
}

func NewRegistry(_ config.Config, clientRegistry *client.Registry) *Registry {
	return &Registry{
		UserInfoBuilder: NewUsersInfoBuilder(clientRegistry.UsersDB),
		UserAuthBuilder: NewUserAuthBuilder(clientRegistry.Twilio),
	}
}
