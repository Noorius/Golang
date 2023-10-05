package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *Application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.HealthCheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/knives", app.createKnifeHandler)
	router.HandlerFunc(http.MethodGet, "/v1/knives/:id", app.showKnifeHandler)

	return router
}