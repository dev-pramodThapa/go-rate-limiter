package tests

import (
	"go-rate-limiter/limiter"
	"go-rate-limiter/routes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRateLimitMiddleware(t *testing.T) {
	// Initialize the rate limiter config
	appConfig := limiter.InitConfig()
	limiter.AppConfig = appConfig

	// Create a new server mux and setup routes
	mux := http.NewServeMux()
	routes.SetupRoutes(mux)

	// Create a test server
	server := httptest.NewServer(mux)
	defer server.Close()

	client := &http.Client{}

	// Simulate sending 5 requests, which is within the default rate limit for users
	for i := 1; i <= 5; i++ {
		req, _ := http.NewRequest(http.MethodGet, server.URL+"/users/testUser", nil)
		req.Header.Set("X-User-ID", "testUser")
		resp, err := client.Do(req)

		if err != nil {
			t.Fatalf("Request %d failed: %v", i, err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected request %d to succeed, got status code %d", i, resp.StatusCode)
		}

		resp.Body.Close()
	}

	// Simulate the 6th request, which should be rate-limited
	req, _ := http.NewRequest(http.MethodGet, server.URL+"/users/testUser", nil)
	req.Header.Set("X-User-ID", "testUser")
	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf("6th request failed: %v", err)
	}

	if resp.StatusCode != http.StatusTooManyRequests {
		t.Fatalf("Expected rate limit to be hit, got status code %d", resp.StatusCode)
	}

	resp.Body.Close()
}
