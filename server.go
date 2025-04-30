package main

import (
	"fmt"
	"net/http"
)

// SetupRoutes configures all the HTTP routes for our server
func SetupRoutes() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/key", apiKeyHandler)
	http.HandleFunc("/simulator/load-balancer", loadBalancerSimulatorHandler)
	http.HandleFunc("/simulator/load-balancer-concurrent", loadBalancerSimulatorConcurrentHandler)
}

// StartHTTPServer starts the HTTP server on the specified address
func StartHTTPServer(addr string) error {
	// Print server info
	fmt.Printf("Server running at http://%s\n", addr)
	fmt.Println("Try these endpoints:")
	fmt.Printf("  - http://%s/health\n", addr)
	fmt.Printf("  - http://%s/api/key\n", addr)
	fmt.Printf("  - http://%s/simulator/load-balancer\n", addr)
	fmt.Println("Press Ctrl+C to stop the server")

	// Start the server
	return http.ListenAndServe(addr, nil)
}
