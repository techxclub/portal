package router

import (
	"github.com/gorilla/mux"
	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/handler"
	"github.com/techx/portal/service"
)

func addAuthRoutes(router *mux.Router, cfg *config.Config, sr *service.Registry) {
	authRouter := router.PathPrefix("/public/auth").Subrouter()

	//	swagger:route POST /public/auth/otp/generate public generateOTP
	//	Responses:
	//		200: GenerateOTPResponse
	//		401:
	// 		400: ErrorResponse
	//		422: ErrorResponse
	//		500: ErrorResponse
	//		503: ErrorResponse
	authRouter.
		Methods(constants.MethodPost).
		Name(constants.APINameGenerateOTP).
		Path("/otp/generate").
		Handler(handler.GenerateOTPHandler(cfg, sr))

	//	swagger:route POST /public/auth/otp/verify public verifyOTP
	//	Responses:
	//		200: VerifyOTPResponse
	//		401:
	// 		400: ErrorResponse
	//		422: ErrorResponse
	//		500: ErrorResponse
	//		503: ErrorResponse
	authRouter.
		Methods(constants.MethodPost).
		Name(constants.APINameVerifyOTP).
		Path("/otp/verify").
		Handler(handler.VerifyOTPHandler(cfg, sr))

	//	swagger:route POST /public/auth/otp/resend public ResendOTP
	//	Responses:
	//		200: VerifyOTPResponse
	//		401:
	// 		400: ErrorResponse
	//		422: ErrorResponse
	//		500: ErrorResponse
	//		503: ErrorResponse
	authRouter.
		Methods(constants.MethodPost).
		Name(constants.APINameResendOTP).
		Path("/otp/resend").
		Handler(handler.ResendOTPHandler(cfg, sr))
}
