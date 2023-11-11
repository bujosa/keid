package main

import (
	"fmt"
	"keid/config"
	"net/http"
)

func main() {
	r := config.NewRouter()

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.ListenAndServe(":8080", nil)
}
