package main

import "net/http"

func (app *application) healthHandler(w http.ResponseWriter, r *http.Request) {
	app.jsonOuput(w, "OK", http.StatusOK, nil)
}
