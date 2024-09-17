package server

import (
	"go-rate-limiter/handlers"
	"go-rate-limiter/routes"
	"go-rate-limiter/utils"
	"net/http"
)

// StartServer initializes and starts the HTTP server
func StartServer() {
	// Initialize the logger
	logger := utils.NewLogger()

	// Create a new HTTP ServeMux
	mux := http.NewServeMux()

	// Setup routes
	routes.SetupRoutes(mux)

	// Apply request logging middleware
	loggedMux := handlers.LogRequestMiddleware(logger)(mux)

	// Define the server port
	port := ":8000"

	// Create the server
	server := &http.Server{
		Addr:    port,
		Handler: loggedMux,
	}

	logger.Info("Server is starting and listening on port " + port)

	// Start the server and log any errors
	if err := server.ListenAndServe(); err != nil {
		logger.Error("Failed to start server: " + err.Error())
	}
}
