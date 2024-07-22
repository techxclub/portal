package middleware

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/techx/portal/client/ratelimiter"
	"github.com/techx/portal/config"
	"github.com/techx/portal/constants"
	"github.com/techx/portal/utils"
)

func RateLimiter(cfg *config.Config, rateLimiterClient ratelimiter.RateLimiter) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !cfg.RateLimitEnabled {
				next.ServeHTTP(w, r)
				return
			}

			apiName := mux.CurrentRoute(r).GetName()
			rateLimitConfig := cfg.RateLimit.GetAPIRateLimitConfig(apiName)
			if !rateLimitConfig.Enabled {
				next.ServeHTTP(w, r)
				return
			}

			uniqueID := utils.GetUniqueRequestID(r)
			if uniqueID == "" {
				uniqueID = constants.GlobalRateLimitKey
			}

			key := rateLimitKey(apiName, uniqueID)
			isAcquired, err := rateLimiterClient.TryAcquire(r.Context(), key, rateLimitConfig.Attempts, rateLimitConfig.WindowSecs)
			if err != nil {
				log.Error().Err(err).Msg(fmt.Sprintf("TryAcquire failed for key: %s ", key))
			}

			if !isAcquired {
				log.Error().Msg(fmt.Sprintf("API request rate limited for path %s for unique id %s", r.URL.Path, uniqueID))
				w.WriteHeader(http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func rateLimitKey(apiName, id string) string {
	return fmt.Sprintf("rate-limit:%s:%s", apiName, id)
}
