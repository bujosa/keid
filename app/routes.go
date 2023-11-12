package app

import (
	"keid/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})

	router.Route(("/users"), loadUserRoutes)

	return router
}

func loadUserRoutes(router chi.Router) {
	userHandler := &handler.User{}

	router.Get("/", userHandler.GetAll)
	router.Post("/", userHandler.Create)
	router.Get("/{id}", userHandler.GetById)
	router.Put("/{id}", userHandler.Update)
	router.Delete("/{id}", userHandler.Delete)
}
