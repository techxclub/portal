package router

import (
	"github.com/gorilla/mux"
	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/handler"
	"github.com/techx/portal/service"
)

func addOAuthRoutes(router *mux.Router, cfg *config.Config, sr *service.Registry) {
	authRouter := router.PathPrefix("/public/google/oauth").Subrouter()

	//	swagger:route GET /public/google/oauth/debug generateOTP
	//	Responses:
	//		200: GoogleOAuthLoginResponse
	//		401:
	// 		400: ErrorResponse
	//		422: ErrorResponse
	//		500: ErrorResponse
	//		503: ErrorResponse
	authRouter.
		Methods(constants.MethodGet).
		Name(constants.APINameGoogleOAuthDebug).
		Path("/debug").
		Handler(handler.GoogleOAuthDebugHandler(cfg))

	//	swagger:route GET /public/google/oauth/login generateOTP
	//	Responses:
	//		200: GoogleOAuthLoginResponse
	//		401:
	// 		400: ErrorResponse
	//		422: ErrorResponse
	//		500: ErrorResponse
	//		503: ErrorResponse
	authRouter.
		Methods(constants.MethodGet).
		Name(constants.APINameGoogleOAuthLogin).
		Path("/login").
		Handler(handler.GoogleOAuthLoginHandler(cfg, sr))

	//	swagger:route GET /public/google/oauth/callback generateOTP
	//	Responses:
	//		200: GoogleOAuthCallbackResponse
	//		401:
	// 		400: ErrorResponse
	//		422: ErrorResponse
	//		500: ErrorResponse
	//		503: ErrorResponse
	authRouter.
		Methods(constants.MethodGet).
		Name(constants.APINameGoogleOAuthCallback).
		Path("/callback").
		Handler(handler.GoogleOAuthCallbackHandler(cfg, sr))

	//	swagger:route POST /public/google/oauth/exchange oauthExchangeCode
	//	Responses:
	//		200: GoogleOAuthExchangeResponse
	//		401:
	// 		400: ErrorResponse
	//		422: ErrorResponse
	//		500: ErrorResponse
	//		503: ErrorResponse
	authRouter.
		Methods(constants.MethodPost).
		Name(constants.APINameGoogleOAuthExchange).
		Path("/exchange").
		Handler(handler.GoogleOAuthExchangeHandler(cfg, sr))
}
