package router

import (
	"github.com/gorilla/mux"
	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/handler"
	"github.com/techx/portal/middleware"
	"github.com/techx/portal/service"
)

func addPublicRoutes(router *mux.Router, cfg *config.Config, sr *service.Registry) {
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
		Methods(constants.MethodPost).
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
		Methods(constants.MethodGet).
		Path("/user/profile").
		Handler(middleware.Authorization(cfg.Auth)(handler.UserProfileHandler(cfg, sr)))

	//	swagger:route GET /public/company/list public companyList
	//	Responses:
	//		200: CompanyListResponse
	//		401:
	// 		400: ErrorResponse
	//		422: ErrorResponse
	//		500: ErrorResponse
	//		503: ErrorResponse
	publicRouter.
		Methods(constants.MethodGet).
		Path("/company/list").
		Handler(handler.CompanyListHandler(cfg, sr))

	//	swagger:route GET /public/company/users/list public companyUsersList
	//	Responses:
	//		200: CompanyUsersListResponse
	//		401:
	// 		400: ErrorResponse
	//		422: ErrorResponse
	//		500: ErrorResponse
	//		503: ErrorResponse
	publicRouter.
		Methods(constants.MethodGet).
		Path("/company/users/list").
		Handler(middleware.Authorization(cfg.Auth)(handler.CompanyUsersListHandler(cfg, sr)))

	//	swagger:route Post /public/user/referral/request public referralRequest
	//	Responses:
	//		200: ReferralResponse
	//		401:
	// 		400: ErrorResponse
	//		422: ErrorResponse
	//		500: ErrorResponse
	//		503: ErrorResponse
	publicRouter.
		Methods(constants.MethodGet).
		Path("/user/referral/request").
		Handler(middleware.Authorization(cfg.Auth)(handler.ReferralHandler(cfg, sr)))
}
