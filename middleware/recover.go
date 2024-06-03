package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

// RecoverMiddleware catches panics and prevents the request goroutine from crashing.
func RecoverMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					path, _ := mux.CurrentRoute(r).GetPathTemplate()
					if path == "" {
						path = r.URL.String()
					}

					log.Error().Msgf("Recovered from panic: %+v\npath: %s\n%s", err, path, string(debug.Stack()))
					w.WriteHeader(http.StatusInternalServerError)

					return
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
