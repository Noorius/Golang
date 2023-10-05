package main

import (
	"fmt"
	"net/http"
)

func (app *Application) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	js := `{"status" : "available", "environment" : %q, "version": %q}`
	js = fmt.Sprintf(js, app.config.env, version)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(js))
}

func (app *Application) CountHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have been here: %d times\n ", cnt)
	cnt += 1
}
