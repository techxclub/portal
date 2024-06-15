package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
)

// AuthVerifier catches panics and prevents the request goroutine from crashing.
func AuthVerifier(authConfig *config.Auth) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if authConfig.DebugMode {
				next.ServeHTTP(w, r)
				return
			}

			authCookie, err := r.Cookie(constants.CookieAuthToken)
			if err != nil {
				http.Error(w, "Missing authorization cookie", http.StatusBadRequest)
				return
			}

			tokenStr := authCookie.Value
			if tokenStr == "" {
				http.Error(w, "Missing token in cookie", http.StatusUnauthorized)
				return
			}

			userID := r.Header.Get(constants.HeaderXUserID)
			if userID == "" {
				http.Error(w, "Missing user id header", http.StatusBadRequest)
				return
			}

			if err := domain.VerifyToken(tokenStr, userID, authConfig); err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
