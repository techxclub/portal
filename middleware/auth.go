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
				http.Error(w, "Missing authorization cookie", http.StatusUnauthorized)
				return
			}

			tokenStr := authCookie.Value
			if tokenStr == "" {
				http.Error(w, "Missing token in cookie", http.StatusUnauthorized)
				return
			}

			if err := domain.VerifyToken(tokenStr, authConfig); err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
