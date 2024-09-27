package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) registerRoutes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/health", app.healthHandler)
	router.HandlerFunc(http.MethodGet, "/v1/users/:id", app.getUser)
	router.HandlerFunc(http.MethodPost, "/v1/users", app.createUser)
	router.HandlerFunc(http.MethodPut, "/v1/users/:id", app.updateUser)

	return router
}
