package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(basicHandler),
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to echo server", err)
	}
}

// Create a handler for your CRUD
func basicHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		
	case http.MethodPost:
	case http.MethodPatch:
	case http.MethodDelete:
	}
}
