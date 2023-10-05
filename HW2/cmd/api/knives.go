package main

import (
	"fmt"
	"hw2.nur.net/internal/data"
	"hw2.nur.net/internal/validator"
	"net/http"
	"time"
)

func (app *Application) createKnifeHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title    string        `json:"title"`
		Material string        `json:"material"`
		Color    string        `json:"color"`
		Country  string        `json:"country"`
		Duration data.Duration `json:"duration"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	knife := &data.Knife{
		Title:    input.Title,
		Material: input.Material,
		Color:    input.Color,
		Country:  input.Country,
		Duration: input.Duration,
	}

	v := validator.New()

	if data.ValidateKnife(v, knife); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
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
