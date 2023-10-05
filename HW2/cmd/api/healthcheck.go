package main

import (
	"fmt"
	"net/http"
)

func (app *Application) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}

func (app *Application) CountHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have been here: %d times\n ", cnt)
	cnt += 1
}
