package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

type ServerManager struct {
	servers    map[int]*http.Server
	mu         sync.Mutex
	portNumber int
}

func NewServerManager(startPort int) *ServerManager {
	return &ServerManager{
		servers:    make(map[int]*http.Server),
		portNumber: startPort,
	}
}

func (serv *ServerManager) StartServers() {
	serv.mu.Lock()
	defer serv.mu.Unlock()
	// concurrenctly do this
	port := serv.portNumber
	// establish that the port number is now in use
	// increment by 1 to prevent port collisions
	serv.portNumber++
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "this is server ", port, " reporting for duty")
	})
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println("server failed for this reason:", err)
		}
	}()
	// add the server to the map
	serv.servers[port] = server
}

func (serv *ServerManager) StopServers() {
	serv.mu.Lock()
	defer serv.mu.Unlock()
	for _, server := range serv.servers {
		server.Shutdown(context.Background())
	}
}
