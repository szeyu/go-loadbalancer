package main

// create a load balancer object
type LoadBalancer struct {
	apiKeys []string // list of api keys
	index   int      // index of the current api key
}

// constructor for the load balancer
// take a list of api keys and return a LoadBalancer object
func NewLoadBalancer(apiKeys []string) LoadBalancer {
	return LoadBalancer{
		apiKeys: apiKeys,
		index:   0, // initalize the first api key
	}
}

// get the next api key
// lb is a LoadBalancer object alias
func (lb *LoadBalancer) GetNextApiKey() string {
	apiKey := lb.apiKeys[lb.index] // get the current api key
	lb.index = (lb.index + 1) % len(lb.apiKeys)

	return apiKey
}
