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
		Name(constants.APINameGoogleSignIn).
		Path("/exchange").
		Handler(handler.GoogleSignInHandler(cfg, sr))
}
