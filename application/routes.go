package application

import (
	"net/http"

	"github.com/emmanuelalo52/order-api/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

	})
	router.Route("/orders", loadOrderRoutes)
	return router
}

func loadOrderRoutes(router chi.Router) {
	orderHandler := &handler.Order{}
	router.Get("/",orderHandler.Create)
	router.Post("/", orderHandler.List)
	router.Get("/{id}",orderHandler.GetbyID)
	router.Put("/{id}",orderHandler.UpdatebyID)
	router.Delete("/{id}", orderHandler.DeletebyID)
}
