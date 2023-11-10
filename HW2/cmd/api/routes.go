package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *Application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.HealthCheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/knives", app.requirePermission("knives:read", app.listKnivesHandler))
	router.HandlerFunc(http.MethodPost, "/v1/knives", app.requirePermission("knives:write", app.createKnifeHandler))
	router.HandlerFunc(http.MethodGet, "/v1/knives/:id", app.requirePermission("knives:read", app.showKnifeHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/knives/:id", app.requirePermission("knives:write", app.updateKnifeHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/knives/:id", app.requirePermission("knives:write", app.deleteKnifeHandler))

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)

	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	return app.recoverPanic(app.rateLimit(app.authenticate(router)))
}
