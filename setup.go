package main

import (
	"github.com/gin-gonic/gin"
)

func setup() {
	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// Initialize necessary components
	InitializeLoadBalancer()
	InitializeServerAddress()

	SetupRoutes()
	StartHTTPServer(serverAddr)
}
