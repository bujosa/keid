package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	db     *redis.Client
}

func New() *App {
	return &App{
		router: loadRoutes(),
		db:     redis.NewClient(&redis.Options{}),
	}
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: a.router,
	}

	err := a.db.Ping(ctx).Err()

	if err != nil {
		return fmt.Errorf("failed to ping redis: %w", err)
	}

	err = server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to listen and serve: %w", err)
	}

	return nil
}
