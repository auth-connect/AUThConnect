package main

import "net/http"

func (app *application) health(w http.ResponseWriter, r *http.Request) {
	message := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
		},
	}

	err := app.jsonWrite(w, message, http.StatusOK, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
