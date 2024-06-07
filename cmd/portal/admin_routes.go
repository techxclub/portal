package main

import (
	"github.com/gorilla/mux"
	"github.com/techx/portal/config"
	"github.com/techx/portal/handler"
	"github.com/techx/portal/service"
)

func addAdminRoutes(router *mux.Router, cfg config.Config, sr *service.Registry) {
	adminRouter := router.PathPrefix("/admin").Subrouter()

	//	swagger:route GET /admin/user/details admin bulkGetUsers
	//	Responses:
	//		200: BulkUserDetailsResponse
	//		401:
	// 		400: ErrorResponse
	//		422: ErrorResponse
	//		500: ErrorResponse
	//		503: ErrorResponse
	adminRouter.
		Methods("GET").
		Path("/user/details").
		Handler(handler.AdminUserDetailsHandler(cfg, sr))
}
