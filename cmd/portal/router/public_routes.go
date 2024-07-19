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
	publicRouter.Use(middleware.Authorization(cfg.Auth))

	//	swagger:route POST /public/user/register public registerUserV1
	//	Responses:
	//		200: RegisterUserV1Response
	//		401:
	// 		400: ErrorResponse
	//		422: ErrorResponse
	//		500: ErrorResponse
	//		503: ErrorResponse
	publicRouter.
		Methods(constants.MethodPost).
		Name(constants.APINameUserRegister).
		Path("/user/register").
		Handler(handler.RegisterUserV1Handler(cfg, sr))

	//	swagger:route POST /public/mentor/register public registerUserV1
	//	Responses:
	//		200: RegisterUserV1Response
	//		401:
	// 		400: ErrorResponse
	//		422: ErrorResponse
	//		500: ErrorResponse
	//		503: ErrorResponse
	publicRouter.
		Methods(constants.MethodPost).
		Name(constants.APINameMentorRegister).
		Path("/mentor/register").
		Handler(handler.RegisterMentorHandler(cfg, sr))

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
		Name(constants.APINameUserProfile).
		Path("/user/profile").
		Handler(handler.UserProfileHandler(cfg, sr))

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
		Name(constants.APINameCompanyList).
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
		Name(constants.APINameCompanyUserList).
		Path("/company/users/list").
		Handler(handler.CompanyUsersListHandler(cfg, sr))

	//	swagger:route GET /public/mentors/list public companyUsersList
	//	Responses:
	//		200: CompanyUsersListResponse
	//		401:
	// 		400: ErrorResponse
	//		422: ErrorResponse
	//		500: ErrorResponse
	//		503: ErrorResponse
	publicRouter.
		Methods(constants.MethodGet).
		Name(constants.APINameMentorList).
		Path("/mentors/list").
		Handler(handler.MentorsListHandler(cfg, sr))

	//	swagger:route Post /public/user/referral/request public referralRequest
	//	Responses:
	//		200: ReferralResponse
	//		401:
	// 		400: ErrorResponse
	//		422: ErrorResponse
	//		500: ErrorResponse
	//		503: ErrorResponse
	publicRouter.
		Methods(constants.MethodPost).
		Name(constants.APINameReferralRequest).
		Path("/user/referral/request").
		Handler(handler.ReferralHandler(cfg, sr))
}
