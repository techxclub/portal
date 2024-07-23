package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
)

func AdminAuth(cfg *config.Config) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		allowedClientID := cfg.Admin.ClientID
		allowedPassKey := cfg.Admin.PassKey

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqClientID := r.Header.Get(constants.HeaderClientID)
			reqPassKey := r.Header.Get(constants.HeaderPasskey)

			if reqClientID != allowedClientID || reqPassKey != allowedPassKey {
				http.Error(w, "invalid client id or pass key for admin", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
