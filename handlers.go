package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Health check handler
func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "ok", "message": "Server is healthy"}
	json.NewEncoder(w).Encode(response)
}

// Handler to demonstrate load balancer functionality
func apiKeyHandler(w http.ResponseWriter, r *http.Request) {
	// Get the next API key from the load balancer
	apiKey := loadBalancer.GetNextApiKey()
	fmt.Println("API Key:", apiKey)

	response := map[string]string{
		"api_key": apiKey,
		"message": "This API key would be used for external service authentication",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func loadBalancerSimulatorHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the number parameter from the URL path
	numStr := r.URL.Query().Get("n")
	if numStr == "" {
		numStr = "10" // Default value if not provided
	}

	// Convert string to int
	num, err := strconv.Atoi(numStr)
	if err != nil {
		http.Error(w, "Invalid number parameter", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Running load balancer simulation"}
	json.NewEncoder(w).Encode(response)

	// Run the simulation with the provided number
	go RunLoadBalancerSimulation(num)
}

func loadBalancerSimulatorConcurrentHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the number parameter from the URL path
	numStr := r.URL.Query().Get("n")
	if numStr == "" {
		numStr = "10" // Default value if not provided
	}

	// Convert string to int
	num, err := strconv.Atoi(numStr)
	if err != nil {
		http.Error(w, "Invalid number parameter", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Running concurrent load balancer simulation"}
	json.NewEncoder(w).Encode(response)

	// We need to implement this function in simulator.go
	go RunLoadBalancerSimulationConcurrent(num)
}
