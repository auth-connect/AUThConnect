package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) registerRoutes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/health", app.health)
	router.HandlerFunc(http.MethodPost, "/v1/users/register", app.registerUser)
	router.HandlerFunc(http.MethodPut, "/v1/users/activate", app.activateUser)
	router.HandlerFunc(http.MethodPost, "/v1/users/login", app.loginUser)

	return app.recoverPanic(app.enableCORS(router))
}
