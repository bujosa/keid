package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type User struct{}

func (u *User) GetAll(w http.ResponseWriter, r *http.Request) {
	// do something
}

func (u *User) GetById(w http.ResponseWriter, r *http.Request) {
	// do something
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	// do something
}

func (u *User) Update(w http.ResponseWriter, r *http.Request) {
	// do something
}

func (u *User) Delete(w http.ResponseWriter, r *http.Request) {
	// do something
}

func LoadUserRoutes(router chi.Router) {
	userHandler := &User{}

	router.Get("/", userHandler.GetAll)
	router.Post("/", userHandler.Create)
	router.Get("/{id}", userHandler.GetById)
	router.Put("/{id}", userHandler.Update)
	router.Delete("/{id}", userHandler.Delete)
}
