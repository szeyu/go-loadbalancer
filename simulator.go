package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Simple structure to represent a service
type Service struct {
	Name string
	ID   int
}

// RunLoadBalancerSimulation starts a simple simulation to demonstrate the load balancer
func RunLoadBalancerSimulation(n int) {
	fmt.Println("\n=== Starting Simple Load Balancer Simulation ===")

	// Create a few services
	services := []Service{
		{Name: "UserAuth", ID: 1},
		{Name: "DataSync", ID: 2},
		{Name: "Analytics", ID: 3},
		{Name: "Messaging", ID: 4},
	}

	// Timing variables
	tickInterval := 500 * time.Millisecond

	// Simulation loop - run until duration is reached
	for i := 0; i < n; i++ {
		// Choose a service (round-robin through services)
		service := services[i%len(services)]

		// Get API key from load balancer
		apiKey := loadBalancer.GetNextApiKey()

		// Print the result
		fmt.Printf("[%s] Service: %s (ID: %d) using API key: %s\n",
			time.Now().Format("15:04:05"), service.Name, service.ID, apiKey)

		// Wait for the next tick
		time.Sleep(tickInterval)
	}

	fmt.Println("=== Simulation Complete ===")
}

// RunLoadBalancerSimulationConcurrent starts a concurrent simulation to demonstrate the load balancer
func RunLoadBalancerSimulationConcurrent(n int) {
	fmt.Println("\n=== Starting Concurrent Load Balancer Simulation ===")

	// Create a few services
	services := []Service{
		{Name: "UserAuth", ID: 1},
		{Name: "DataSync", ID: 2},
		{Name: "Analytics", ID: 3},
		{Name: "Messaging", ID: 4},
	}

	// Use a wait group to track goroutines
	var wg sync.WaitGroup
	wg.Add(n)

	// Launch n concurrent requests
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()

			// Choose a service (round-robin through services)
			service := services[i%len(services)]

			// Get API key from load balancer
			apiKey := loadBalancer.GetNextApiKey()

			// Print the result
			fmt.Printf("[%s] Concurrent - Service: %s (ID: %d) using API key: %s\n",
				time.Now().Format("15:04:05"), service.Name, service.ID, apiKey)

			// Add a small random delay to simulate varying process times
			time.Sleep(time.Duration(100+rand.Intn(500)) * time.Millisecond)
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	fmt.Println("=== Concurrent Simulation Complete ===")
}
