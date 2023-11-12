package app

import "net/http"

type App struct {
	router http.Handler
}

func NewApp() *App {
	return &App{
		router: loadRoutes(),
	}
}
