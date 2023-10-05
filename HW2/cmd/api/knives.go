package main

import (
	"fmt"
	"net/http"
)

func (app *Application) createKnifeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "created a new knife")
}

func (app *Application) showKnifeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "showing the detail of a knife with %d\n", id)
}
