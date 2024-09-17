package routes

import (
	"go-rate-limiter/handlers"
	"net/http"
)

func SetupRoutes(mux *http.ServeMux) {
	mux.Handle("/users/", handlers.RateLimitMiddleware(http.HandlerFunc(handlers.UserDataHandler)))
	mux.Handle("/admin/", handlers.RateLimitMiddleware(http.HandlerFunc(handlers.AdminDashboardHandler)))

	mux.HandleFunc("/metrics", handlers.MetricsHandler)

	mux.HandleFunc("/public/info", handlers.PublicInfoHandler)

	mux.HandleFunc("/update-rate-limit", handlers.UpdateRateLimitConfigHandler)
}
