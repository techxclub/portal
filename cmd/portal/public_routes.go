package main

import (
	"github.com/gorilla/mux"
	"github.com/techx/portal/config"
	"github.com/techx/portal/handler"
	"github.com/techx/portal/service"
)

func addPublicRoutes(router *mux.Router, cfg config.Config, sr *service.Registry) {
	publicRouter := router.PathPrefix("/public").Subrouter()

	//	swagger:route POST /public/user/register/v1 public registerUserV1
	//	Responses:
	//		200: RegisterUserV1Response
	//		401:
	// 		400: ErrorResponse
	//		422: ErrorResponse
	//		500: ErrorResponse
	//		503: ErrorResponse
	publicRouter.
		Methods("POST").
		Path("/user/register/v1").
		Handler(handler.RegisterUserV1Handler(cfg, sr))

	//	swagger:route GET /public/user/profile public userProfile
	//	Responses:
	//		200: UserProfile
	//		401:
	// 		400: ErrorResponse
	//		422: ErrorResponse
	//		500: ErrorResponse
	//		503: ErrorResponse
	publicRouter.
		Methods("GET").
		Path("/public/user/profile").
		Handler(handler.UserProfileHandler(cfg, sr))
}
