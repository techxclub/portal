package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/techx/portal/config"
	"github.com/techx/portal/service"
)

func NewRouter(cfg *config.Config, sr *service.Registry) *mux.Router {
	router := mux.NewRouter()

	// Enable Swagger if enabled
	router.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", http.FileServer(
		http.Dir(cfg.Swagger.Path))))

	addPublicRoutes(router, *cfg, sr)
	return router
}
