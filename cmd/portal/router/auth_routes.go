package router

import (
	"github.com/gorilla/mux"
	"github.com/techx/portal/config"
	"github.com/techx/portal/handler"
	"github.com/techx/portal/service"
)

func addAuthRoutes(router *mux.Router, cfg config.Config, sr *service.Registry) {
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
		Methods("POST").
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
		Methods("POST").
		Path("/otp/verify").
		Handler(handler.VerifyOTPHandler(cfg, sr))
}
