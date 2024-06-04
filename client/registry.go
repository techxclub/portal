package client

import (
	"github.com/techx/portal/client/db"
	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
)

type Registry struct {
	UsersDB *db.Repository
}

func NewRegistry(cfg config.Config) *Registry {
	usersDB, err := db.NewRepository(cfg, constants.TableNameUsers)
	if err != nil {
		panic(err)
	}

	return &Registry{
		UsersDB: usersDB,
	}
}
