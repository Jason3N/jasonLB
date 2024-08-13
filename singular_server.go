package main
// this is to demonstarte the use of singular server using http


import (
	"fmt"
    "net/http"
)

func main() {
	// create a handler that returns a response of hello
	myHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "this is server 1 reporting for duty")
    })
    s1 := &http.Server{
        Addr:           ":8081",
        Handler:        myHandler,
    }

	fmt.Println("backend server are up")

    if err := s1.ListenAndServe(); err != nil {
        fmt.Println("Server failed:", err)
    }
}