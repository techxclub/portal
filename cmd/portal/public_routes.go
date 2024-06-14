package main

import (
	"github.com/gorilla/mux"
	"github.com/techx/portal/config"
	"github.com/techx/portal/handler"
	"github.com/techx/portal/middleware"
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
		Path("/user/profile").
		Handler(middleware.AuthVerifier(cfg.Auth)(handler.UserProfileHandler(cfg, sr)))

	//	swagger:route Post /public/user/referral/request public userProfile
	//	Responses:
	//		200: UserProfile
	//		401:
	// 		400: ErrorResponse
	//		422: ErrorResponse
	//		500: ErrorResponse
	//		503: ErrorResponse
	publicRouter.
		Methods("POST").
		Path("/user/referral/request").
		Handler(middleware.AuthVerifier(cfg.Auth)(handler.ReferralHandler(cfg, sr)))

	//	swagger:route GET /public/company/list public companyList
	//	Responses:
	//		200: CompanyList
	//		401:
	// 		400: ErrorResponse
	//		422: ErrorResponse
	//		500: ErrorResponse
	//		503: ErrorResponse
	publicRouter.
		Methods("GET").
		Path("/company/list").
		Handler(handler.CompanyListHandler(cfg, sr))

	//	swagger:route GET /public/company/users/list public companyUsersList
	//	Responses:
	//		200: CompanyUsersList
	//		401:
	// 		400: ErrorResponse
	//		422: ErrorResponse
	//		500: ErrorResponse
	//		503: ErrorResponse
	publicRouter.
		Methods("GET").
		Path("/company/users/list").
		Handler(middleware.AuthVerifier(cfg.Auth)(handler.CompanyUsersListHandler(cfg, sr)))
}
