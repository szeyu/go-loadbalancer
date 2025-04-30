package main

func setup() {
	InitializeLoadBalancer()
	InitializeServerAddress()

	SetupRoutes()
	StartHTTPServer(serverAddr)
}
