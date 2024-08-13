package main
// this is to demonstarte the use of multiple servers

import (
	"fmt"
    "net/http"
)

func main() {
	// create two handlers that returns a response of hello
	handler1 := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "this is server 1 reporting for duty")
	})
	handler2 := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "this is server 2 reporting for duty")
	})
	// two servers
	s1 := &http.Server{
		Addr:		   ":8081",
		Handler:	   handler1,
	}
	s2 := &http.Server{
		Addr:		   ":8082",
		Handler:	   handler2,
	}

	fmt.Println("backend servers are up!")

	// use goroutine to run the servers concurrently
	go func () {
		if err := s1.ListenAndServe(); err != nil {
			fmt.Println("Server failed:", err)
		}
	}()

	go func () {
		if err := s2.ListenAndServe(); err != nil {
			fmt.Println("Server failed:", err)
		}	
	}()

	select {}
}