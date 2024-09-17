package handlers

import (
	"go-rate-limiter/limiter"
	"go-rate-limiter/utils"
	"net/http"
)

// LogRequestMiddleware logs incoming requests and their responses
func LogRequestMiddleware(logger *utils.CustomLogger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Request(r.Method, r.URL.Path)
			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}
}

// RateLimitMiddleware applies the rate limiting logic based on endpoint type and ID
func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		segments := utils.SplitPath(path)
		if len(segments) < 3 || utils.IsStringEmpty(segments[2]) {
			http.Error(w, "Invalid request.", http.StatusBadRequest)
			return
		}

		endpointType := segments[1]

		// Get ID from request header, fallback to path if not found
		id := r.Header.Get("X-User-ID")
		if utils.IsStringEmpty(id) {
			// Fallback if no header ID
			id = segments[2]
		}

		// Get the configuration from the global config
		conf := limiter.GetConfig()

		endpointConfig := conf.GetConfig(endpointType, id)

		bucket := conf.RateLimiter.GetBucket(id, endpointType, endpointConfig.MaxTokens, endpointConfig.RefillRate)

		// Update metrics for total requests count
		conf.Metrics.IncreaseRequestsCount(segments[1] + ":" + id)

		if !bucket.AllowRequest() {
			// Rate limit exceeded
			conf.Metrics.IncreaseThrottledRequest(segments[1] + ":" + id)
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		// Call the next handler if rate limit is not exceeded
		next.ServeHTTP(w, r)
	})
}
