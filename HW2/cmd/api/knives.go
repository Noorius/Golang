package main

import (
	"errors"
	"fmt"
	"hw2.nur.net/internal/data"
	"hw2.nur.net/internal/validator"
	"net/http"
	"strconv"
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

	err = app.models.Knives.Insert(knife)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/movies/%d", knife.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"movie": knife}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *Application) showKnifeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	knives, err := app.models.Knives.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"knife": knives}, nil)
	if err != nil {
		app.logger.PrintError(err, nil)
		app.serverErrorResponse(w, r, err)
	}
}

func (app *Application) updateKnifeHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the movie ID from the URL.
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	// Fetch the existing movie record from the database, sending a 404 Not Found
	// response to the client if we couldn't find a matching record.
	knive, err := app.models.Knives.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return

	}

	if r.Header.Get("X-Expected-Version") != "" {
		if strconv.FormatInt(int64(knive.Version), 32) != r.Header.Get("X-Expected-Version") {
			app.editConflictResponse(w, r)
			return
		}
	}

	// Declare an input struct to hold the expected data from the client.
	var input struct {
		Title    *string        `json:"title"`
		Material *string        `json:"material"`
		Color    *string        `json:"color"`
		Country  *string        `json:"country"`
		Duration *data.Duration `json:"duration"`
	}

	// Read the JSON request body data into the input struct.
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	// Copy the values from the request body to the appropriate fields of the movie
	// record.
	if input.Title != nil {
		knive.Title = *input.Title
	}

	if input.Material != nil {
		knive.Material = *input.Material
	}

	if input.Color != nil {
		knive.Color = *input.Color
	}

	if input.Country != nil {
		knive.Country = *input.Country
	}

	if input.Duration != nil {
		knive.Duration = *input.Duration
	}
	// Validate the updated movie record, sending the client a 422 Unprocessable Entity
	// response if any checks fail.
	v := validator.New()
	if data.ValidateKnife(v, knive); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	// Pass the updated movie record to our new Update() method.
	err = app.models.Knives.Update(knive)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	// Write the updated movie record in a JSON response.
	err = app.writeJSON(w, http.StatusOK, envelope{"knife": knive}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *Application) deleteKnifeHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Knives.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "knife successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *Application) listKnivesHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Title    string
		Material string
		Color    string
		Country  string
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Title = app.readString(qs, "title", "")
	input.Material = app.readString(qs, "material", "")
	input.Color = app.readString(qs, "color", "")
	input.Country = app.readString(qs, "country", "")

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "title", "material", "color", "country", "duration", "-id", "-title", "-material", "-color", "-country", "-duration"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	knives, metadata, err := app.models.Knives.GetAll(input.Title, input.Material, input.Color, input.Country, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// Send a JSON response containing the movie data.
	err = app.writeJSON(w, http.StatusOK, envelope{"knives": knives, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
