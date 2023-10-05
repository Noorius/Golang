package main

import (
	"fmt"
	"net/http"
)

func (app *Application) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {

	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)

	if err != nil {
		app.logger.Println(err)
		http.Error(w, "Server cannot process your request", http.StatusInternalServerError)
	}
}

func (app *Application) CountHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have been here: %d times\n ", cnt)
	cnt += 1
}
