package handlers

import (
	"encoding/json"
	"go-rate-limiter/limiter"
	"net/http"
	"time"
)

func UserDataHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "User data"}
	json.NewEncoder(w).Encode(response)
}

func AdminDashboardHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Admin dashboard"}
	json.NewEncoder(w).Encode(response)
}

func PublicInfoHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Public info available without rate limiting"}
	json.NewEncoder(w).Encode(response)
}

func UpdateRateLimitConfigHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserType   string `json:"user_type"`
		ID         string `json:"id"`
		MaxTokens  int    `json:"max_tokens"`
		RefillRate int    `json:"refill_rate"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.UserType != "users" && req.UserType != "admin" {
		http.Error(w, "Invalid user type", http.StatusBadRequest)
		return
	}

	// Get the configuration from the global config
	config := limiter.GetConfig()

	config.UpdateConfig(req.UserType, req.ID, req.MaxTokens, time.Duration(req.RefillRate)*time.Second)

	// Reset the token bucket for a particular ID and user type
	config.RateLimiter.ResetBucket(req.ID, req.UserType)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Rate limit configuration updated successfully"))
}

// MetricsHandler handles the /metrics endpoint.
func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	metrics := limiter.AppConfig.Metrics.GetMetrics()
	json.NewEncoder(w).Encode(metrics)
}
