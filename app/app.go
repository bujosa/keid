package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	db     *redis.Client
}

func New() *App {
	app := &App{
		db: redis.NewClient(&redis.Options{}),
	}

	app.loadRoutes()
	return app
}

// Start starts the application server.
// It pings the Redis database and listens on port 8080 for incoming requests.
// It returns an error if there was a problem starting the server.
func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: a.router,
	}

	err := a.db.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to ping redis: %w", err)
	}
	defer func() {
		if err := a.db.Close(); err != nil {
			fmt.Printf("failed to close redis connection: %v\n", err)
		}
	}()

	fmt.Println("Starting server on port 8080")

	ch := make(chan error, 1)

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to listen and serve: %w", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}
}
