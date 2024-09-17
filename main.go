package main

import (
	"go-rate-limiter/limiter"
	"go-rate-limiter/server"
)

func main() {
	appConfig := limiter.InitConfig()

	// Set global app config
	limiter.AppConfig = appConfig

	// Start the server
	server.StartServer()
}
