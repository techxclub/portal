package router

import (
	"github.com/gorilla/mux"
	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/handler"
	"github.com/techx/portal/middleware"
	"github.com/techx/portal/service"
)

func addAdminRoutes(router *mux.Router, cfg *config.Config, sr *service.Registry) {
	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.Use(middleware.AdminAuth(cfg))

	//	swagger:route GET /admin/user/list admin getUserList
	//	Responses:
	//		200: UserListResponse
	// 		400: ErrorResponse
	//		500: ErrorResponse
	adminRouter.
		Methods(constants.MethodGet).
		Name(constants.APINameAdminUserList).
		Path("/user/list").
		Handler(handler.AdminUserListHandler(cfg, sr))

	//	swagger:route PUT /admin/user/Update admin updateUserDetails
	//	Responses:
	//		200: SuccessResponse
	// 		400: ErrorResponse
	//		500: ErrorResponse
	adminRouter.
		Methods(constants.MethodPut).
		Name(constants.APINameAdminUserUpdate).
		Path("/user/update").
		Handler(handler.AdminUserUpdateHandler(cfg, sr))

	//	swagger:route GET /admin/company/list admin getCompanyListDetails
	//	Responses:
	//		200: AdminCompanyListResponse
	// 		400: ErrorResponse
	//		500: ErrorResponse
	adminRouter.
		Methods(constants.MethodGet).
		Name(constants.APINameAdminCompanyList).
		Path("/company/list").
		Handler(handler.AdminCompanyListHandler(cfg, sr))

	//	swagger:route GET /admin/referral/list admin getUserReferralList
	//	Responses:
	//		200: AdminUserReferralListResponse
	// 		400: ErrorResponse
	//		500: ErrorResponse
	adminRouter.
		Methods(constants.MethodGet).
		Path("/referral/list").
		Handler(handler.AdminUserReferralListHandler(cfg, sr))

	//	swagger:route PUT /admin/company/Update admin updateCompanyDetails
	//	Responses:
	//		200: SuccessResponse
	// 		400: ErrorResponse
	//		500: ErrorResponse
	adminRouter.
		Methods(constants.MethodPut).
		Name(constants.APINameAdminCompanyUpdate).
		Path("/company/update").
		Handler(handler.AdminCompanyUpdateHandler(cfg, sr))

	//	swagger:route PUT /admin/referral/Update admin updateReferralDetails
	//	Responses:
	//		200: SuccessResponse
	// 		400: ErrorResponse
	//		500: ErrorResponse
	adminRouter.
		Methods(constants.MethodPut).
		Path("/referral/update").
		Handler(handler.AdminReferralUpdateHandler(cfg, sr))
}
