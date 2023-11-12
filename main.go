package main

import (
	"context"
	"keid/app"
	"os"
	"os/signal"
)

func main() {
	app := app.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		panic(err)
	}

}
