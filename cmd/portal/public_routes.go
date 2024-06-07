package main

import (
	"github.com/gorilla/mux"
	"github.com/techx/portal/config"
	"github.com/techx/portal/handler"
	"github.com/techx/portal/service"
)

func addPublicRoutes(router *mux.Router, cfg config.Config, sr *service.Registry) {
	//	swagger:route POST /portal/v1/register public registerUserV1
	//	Responses:
	//		200: RegisterUserV1Response
	//		401:
	// 		400: ErrorResponse
	//		422: ErrorResponse
	//		500: ErrorResponse
	//		503: ErrorResponse
	router.
		Methods("POST").
		Path("/public/user/register/v1").
		Handler(handler.RegisterUserV1Handler(cfg, sr))

	//	swagger:route GET /portal/user/details public userDetails
	//	Responses:
	//		200: UserDetailsResponse
	//		401:
	// 		400: ErrorResponse
	//		422: ErrorResponse
	//		500: ErrorResponse
	//		503: ErrorResponse
	router.
		Methods("GET").
		Path("/public/user/profile").
		Handler(handler.UserProfileHandler(cfg, sr))
}
