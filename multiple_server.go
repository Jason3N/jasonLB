package main
// this is to demonstarte the use of multiple servers

import (
	"fmt"
    "net/http"
)

func main() {
	// what this does is create a server that listens on port 8081
	// and should return a response of hello when curled
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "this is server 1 reporting for duty")
    })
    fmt.Println("backend server is up")
    http.ListenAndServe(":8080", nil)
}