package main

// this is to demonstarte the use of multiple servers

import (
	"fmt"
)

// create a load balancer
type LoadBalancer struct {
	backendServers []*Backend
	current        int
}

// create a new load balancer
func NewLoadBalancer() *LoadBalancer {
	return &LoadBalancer{
		backendServers: make([]*Backend, 0),
		current:        0,
	}
}

// adding backend to the load balancer
func (lb *LoadBalancer) AddBackend(backend *Backend) {
	lb.backendServers = append(lb.backendServers, backend)
}

// round robin algorithm
func (lb *LoadBalancer) RoundRobin() *Backend {
	lb.current = (lb.current + 1) % len(lb.backendServers)
	backend := lb.backendServers[lb.current]
	fmt.Printf("Forwarding request to server at %s\n", backend.URL.String())
	return backend
}
