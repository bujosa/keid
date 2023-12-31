package app

import (
	"keid/controller"
	"keid/repository/user"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route(("/users"), a.LoadUserRoutes)

	a.router = router
}

func (a *App) LoadUserRoutes(router chi.Router) {
	userHandler := &controller.User{
		Repository: &user.UserRepository{
			Client: a.db,
		},
	}

	router.Get("/", userHandler.GetAll)
	router.Post("/", userHandler.Create)
	router.Get("/{id}", userHandler.GetById)
	router.Put("/{id}", userHandler.Update)
	router.Delete("/{id}", userHandler.Delete)
}
