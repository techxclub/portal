package handler

import (
	"encoding/json"
	"net/http"

	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
	"github.com/techx/portal/handler/response"
	"github.com/techx/portal/service"
)

func GoogleSignInHandler(cfg *config.Config, serviceRegistry *service.Registry) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req domain.GoogleOAuthExchangeRequest

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			response.InstrumentErrorResponse(r, errors.AsServiceError(err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		userProfile, err := serviceRegistry.OAuthService.GoogleSignIn(r.Context(), req)
		if err != nil {
			response.InstrumentErrorResponse(r, errors.AsServiceError(err))
			return
		}

		authToken, err := domain.GenerateToken(userProfile.UserUUID, cfg.AuthToken)
		if err != nil {
			response.InstrumentErrorResponse(r, errors.AsServiceError(errors.ErrGeneratingAuthToken))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set(constants.HeaderAuthToken, authToken)
		w.WriteHeader(http.StatusOK)
	}
}
