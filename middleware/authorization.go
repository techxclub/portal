package middleware

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/techx/portal/apicontext"
	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/domain"
)

// Authorization is a middleware that checks if the request is authorized
func Authorization(authConfig config.TokenConfig) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !authConfig.Enabled {
				next.ServeHTTP(w, r)
				return
			}

			authHeader := r.Header.Get(constants.HeaderAuthorization)
			if authHeader == "" {
				http.Error(w, "Missing authorization header", http.StatusBadRequest)
				return
			}

			tokens := strings.Split(authHeader, " ")
			if len(tokens) != 2 || tokens[0] != "Bearer" || tokens[1] == "" {
				http.Error(w, "Invalid authorization header", http.StatusBadRequest)
				return
			}

			authToken := tokens[1]
			userUUID, err := domain.VerifyToken(authConfig, authToken)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			r.Header.Set(constants.HeaderXUserUUID, userUUID)
			reqCtx := apicontext.NewRequestContextFromHTTP(r)
			ctx := apicontext.NewContextWithRequestContext(r.Context(), reqCtx)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
