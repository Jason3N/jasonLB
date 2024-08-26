package main

import (
	"net/http/httputil"
	"net/url"
	"sync"
)

type LeastBackend struct {
	URL               *url.URL
	ReverseProxy      *httputil.ReverseProxy
	activeConnections int
	mu                sync.Mutex
}

func (b *LeastBackend) addConnection() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.activeConnections++
}

func (b *LeastBackend) removeConnection() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.activeConnections--
}

func (b *LeastBackend) GetActiveConnections() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.activeConnections
}

// create a load balancer
type LeastConnectionsLoadBalancer struct {
	backendServers []*Backend
	current        int
}

// create a new load balancer
func NewLeastLoadBalancer() *LeastConnectionsLoadBalancer {
	return &LeastConnectionsLoadBalancer{
		backendServers: make([]*Backend, 0),
		current:        0,
	}
}

// adding backend to the load balancer
func (lb *LeastConnectionsLoadBalancer) AddBackend(backend *Backend) {
	lb.backendServers = append(lb.backendServers, backend)
}

// round robin algorithm
func (lb *LeastConnectionsLoadBalancer) LeastConnections() *Backend {
	return nil
}
