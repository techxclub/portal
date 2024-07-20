package handler

import (
	"net/http"

	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
	"github.com/techx/portal/handler/response"
	"github.com/techx/portal/service"
)

func GoogleOAuthDebugHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		if !cfg.GoogleAuth.Debug {
			http.Error(w, "Debug mode is not enabled", http.StatusForbidden)
			return
		}

		htmlIndex := `
		<html>
		<head>
			<style>
				.container {
					display: flex;
					justify-content: center;
					align-items: center;
					height: 100vh;
				}
				.button {
					background-color: #4285F4;
					color: white;
					padding: 10px 20px;
					text-align: center;
					text-decoration: none;
					display: inline-block;
					font-size: 16px;
					border-radius: 4px;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<a class="button" href="/public/google/oauth/login">Sign in with Google</a>
			</div>
		</body>
		</html>`
		_, err := w.Write([]byte(htmlIndex))
		if err != nil {
			return
		}
	}
}

func GoogleOAuthLoginHandler(_ *config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := serviceRegistry.OAuthService.GoogleLoginURL().RedirectURI
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

func GoogleOAuthCallbackHandler(cfg *config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state := r.FormValue("state")
		code := r.FormValue("code")
		req := domain.GoogleOAuthCallbackRequest{
			State: state,
			Code:  code,
		}

		userProfile, err := serviceRegistry.OAuthService.GoogleOAuthCallback(r.Context(), req)
		if err != nil {
			response.InstrumentErrorResponse(r, errors.AsServiceError(err))
			http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
			return
		}

		authToken, err := domain.GenerateToken(userProfile.UserUUID, cfg.Auth)
		if err != nil {
			response.InstrumentErrorResponse(r, errors.AsServiceError(errors.ErrGeneratingAuthToken))
			w.WriteHeader(http.StatusInternalServerError)
			http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
			return
		}

		w.Header().Set(constants.HeaderAuthToken, authToken)
		w.WriteHeader(http.StatusOK)
		http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
	}
}
