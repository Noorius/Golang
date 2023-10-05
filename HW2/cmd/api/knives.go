package main

import (
	"fmt"
	"hw2.nur.net/internal/data"
	"net/http"
	"time"
)

func (app *Application) createKnifeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "created a new knife")
}

func (app *Application) showKnifeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	knives := data.Knife{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Bark River Classic",
		Material:  "Steel",
		Color:     "Silver",
		Country:   "USA",
		Duration:  2,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"knife": knives}, nil)
	if err != nil {
		app.logger.Println(err)
		app.serverErrorResponse(w, r, err)
	}
}
