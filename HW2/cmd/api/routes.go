package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *Application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/knives", app.listKnivesHandler)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.HealthCheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/knives", app.createKnifeHandler)
	router.HandlerFunc(http.MethodGet, "/v1/knives/:id", app.showKnifeHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/knives/:id", app.updateKnifeHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/knives/:id", app.deleteKnifeHandler)

	return router
}
