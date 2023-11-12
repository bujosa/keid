package handler

import "net/http"

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
