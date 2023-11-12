package main

import (
	"context"
	"keid/app"
)

func main() {
	app := app.New()

	err := app.Start(context.TODO())
	if err != nil {
		panic(err)
	}
}
