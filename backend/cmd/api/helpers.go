package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) jsonOuput(w http.ResponseWriter, payload interface{}, status int, headers http.Header) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// For CLI usage
	body = append(body, '\n')

	for k, v := range headers {
		w.Header()[k] = v
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(body)

	return nil
}
