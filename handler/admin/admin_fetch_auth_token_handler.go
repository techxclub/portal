package admin

import (
	"encoding/json"
	"net/http"

	"github.com/techx/portal/config"
	"github.com/techx/portal/domain"
	"github.com/techx/portal/errors"
	"github.com/techx/portal/handler/response"
)

func FetchAuthTokenHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			UserUUID string `json:"user_id"`
		}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			response.InstrumentErrorResponse(r, errors.AsServiceError(err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		authToken, err := domain.GenerateToken(req.UserUUID, cfg.Auth)
		if err != nil {
			response.InstrumentErrorResponse(r, errors.AsServiceError(errors.ErrGeneratingAuthToken))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp := struct {
			Authorization string `json:"authorization"`
		}{
			Authorization: authToken,
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(resp)
	}
}
