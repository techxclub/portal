package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/techx/portal/builder"
	"github.com/techx/portal/client"
	"github.com/techx/portal/config"
	"github.com/techx/portal/middleware"
	"github.com/techx/portal/service"
)

func NewRouter(cfg *config.Config, cr *client.Registry, br *builder.Registry, sr *service.Registry) *mux.Router {
	router := mux.NewRouter()

	// Enable Swagger if enabled
	router.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", http.FileServer(
		http.Dir(cfg.Swagger.Path))))
	router.Use(middleware.RecoverMiddleware())

	addOAuthRoutes(router, cfg, cr, br, sr)
	addOTPRoutes(router, cfg, cr, br, sr)
	addPublicRoutes(router, cfg, cr, br, sr)
	addAdminRoutes(router, cfg, cr, br, sr)

	return router
}
