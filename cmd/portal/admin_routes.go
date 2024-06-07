package main

import (
	"github.com/gorilla/mux"
	"github.com/techx/portal/config"
	"github.com/techx/portal/handler"
	"github.com/techx/portal/service"
)

func addAdminRoutes(router *mux.Router, cfg config.Config, sr *service.Registry) {
	//	swagger:route GET /admin/user/details admin adminDetails
	//	Responses:
	//		200: AdminDetailsResponse
	//		401:
	// 		400: ErrorResponse
	//		422: ErrorResponse
	//		500: ErrorResponse
	//		503: ErrorResponse
	router.
		Methods("GET").
		Path("/admin/user/details").
		Handler(handler.AdminUserDetailsHandler(cfg, sr))
}
