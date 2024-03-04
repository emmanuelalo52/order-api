package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	router.Get("/test", basicHandler)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to echo server", err)
	}
}

// Create a handler for your CRUD
func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("Testing")))
	// switch r.Method {
	// case http.MethodGet:

	// case http.MethodPost:
	// case http.MethodPatch:
	// case http.MethodDelete:
	// }
}
