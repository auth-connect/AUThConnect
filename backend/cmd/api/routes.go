package main

import "net/http"

func (app *application) registerRoutes() http.Handler {
	mux := http.ServeMux{}

	mux.HandleFunc("/health", app.healthHandler)
	mux.HandleFunc("/", app.getUser)

	return &mux
}
