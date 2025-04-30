package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all the HTTP routes for our server using Gin
func SetupRoutes() *gin.Engine {
	// Create a default gin router with Logger and Recovery middleware already attached
	router := gin.Default()

	// Set up route group for API endpoints
	router.GET("/health", healthHandler)
	router.GET("/api/key", apiKeyHandler)
	router.GET("/simulator/load-balancer", loadBalancerSimulatorHandler)
	router.GET("/simulator/load-balancer-concurrent", loadBalancerSimulatorConcurrentHandler)

	return router
}

// StartHTTPServer starts the HTTP server on the specified address
func StartHTTPServer(addr string) error {
	// Get router from SetupRoutes
	router := SetupRoutes()

	// Print server info
	fmt.Printf("Server running at http://%s\n", addr)
	fmt.Println("Try these endpoints:")
	fmt.Printf("  - http://%s/health\n", addr)
	fmt.Printf("  - http://%s/api/key\n", addr)
	fmt.Printf("  - http://%s/simulator/load-balancer\n", addr)
	fmt.Printf("  - http://%s/simulator/load-balancer-concurrent\n", addr)
	fmt.Println("Press Ctrl+C to stop the server")

	// Start the server
	return router.Run(addr)
}
