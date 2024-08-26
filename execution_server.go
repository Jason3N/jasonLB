package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// this is to demonstarte the use of singular server using http

type Backend struct {
	URL          *url.URL
	ReverseProxy *httputil.ReverseProxy
}

func main() {
	serv := NewServerManager(8081)
	// create 3 servers
	for i := 0; i < 3; i++ {
		serv.StartServers()
	}
	fmt.Println("backend servers are up!")
	// create a load balancer
	lb := NewLoadBalancer()
	// create the backend ports
	backendPorts := []int{8081, 8082, 8083}
	// add the backend ports to the load balancer
	for _, port := range backendPorts {
		url, err := url.Parse(fmt.Sprintf("http://localhost:%d", port))
		if err != nil {
			log.Fatal(err)
		}
		proxy := httputil.NewSingleHostReverseProxy(url)
		backend := &Backend{
			URL:          url,
			ReverseProxy: proxy,
		}
		lb.AddBackend(backend)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		backend := lb.RoundRobin()
		backend.ReverseProxy.ServeHTTP(w, r)
	})

	fmt.Println("Load balancer running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Load balancer failed: %s\n", err)
	}

}
