package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// init runs when the package is initialized
func init() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found or couldn't be loaded")
	}
}

// InitializeServerAddress initializes the server address from environment variables or default
func InitializeServerAddress() {
	// Determine port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if PORT is not set
	}

	serverAddr = fmt.Sprintf("localhost:%s", port)
}

// InitializeLoadBalancer sets up the load balancer with API keys
func InitializeLoadBalancer() {
	// Initialize api keys from env
	apiKeys := os.Getenv("API_KEYS")
	if apiKeys == "" {
		log.Fatal("API_KEYS environment variable is not set")
		apiKeys = "key1,key2,key3,key4,key5" // set default api keys
	}

	apiKeysList := strings.Split(apiKeys, ",")
	loadBalancer = NewLoadBalancer(apiKeysList)

	// Print some examples to show how it works in the console
	fmt.Println("Load Balancer Demo")
	fmt.Println("------------------")
	fmt.Println("Available API Keys:")
	for i := 0; i < 5; i++ {
		fmt.Printf("API Key %d: %s\n", i+1, loadBalancer.GetNextApiKey())
	}
	fmt.Println("------------------")
}
