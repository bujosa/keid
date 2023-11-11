package config

import "github.com/go-chi/chi/v5"

// GetRounter is a function to get router this
// was created for abstraction of router for make
// it easy to change router in the future
func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	return router
}
