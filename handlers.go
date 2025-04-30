package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Health check handler
func healthHandler(c *gin.Context) {
	fmt.Println("Health check received")
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Server is healthy",
	})
}

// Handler to demonstrate load balancer functionality
func apiKeyHandler(c *gin.Context) {
	// Get the next API key from the load balancer
	apiKey := loadBalancer.GetNextApiKey()
	fmt.Println("API Key:", apiKey)

	c.JSON(http.StatusOK, gin.H{
		"api_key": apiKey,
		"message": "This API key would be used for external service authentication",
	})
}

func loadBalancerSimulatorHandler(c *gin.Context) {
	// Extract the number parameter from the URL path
	numStr := c.DefaultQuery("n", "10") // Default value if not provided

	// Convert string to int
	num, err := strconv.Atoi(numStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number parameter"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Running load balancer simulation"})

	// Run the simulation with the provided number
	go RunLoadBalancerSimulation(num)
}

func loadBalancerSimulatorConcurrentHandler(c *gin.Context) {
	// Extract the number parameter from the URL path
	numStr := c.DefaultQuery("n", "10") // Default value if not provided

	// Convert string to int
	num, err := strconv.Atoi(numStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number parameter"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Running concurrent load balancer simulation"})

	// We need to implement this function in simulator.go
	go RunLoadBalancerSimulationConcurrent(num)
}
