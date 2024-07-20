package router

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

	addOAuthRoutes(router, cfg, sr)
	addOTPRoutes(router, cfg, sr)
	addPublicRoutes(router, cfg, sr)
	addAdminRoutes(router, cfg, sr)

	return router
}
