package client

import (
	"github.com/techx/portal/config"
)

type Registry struct{}

func NewRegistry(_ config.Config) *Registry {
	return &Registry{}
}
