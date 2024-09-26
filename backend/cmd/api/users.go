package main

import (
	"net/http"
	"strconv"
	"strings"
)

func (app *application) getUser(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	lastIndex := strings.LastIndex(path, "/")
	if lastIndex == -1 || lastIndex == len(path)-1 {
		app.jsonOuput(w, "LOL", http.StatusBadRequest, nil)
	}

	idStr := path[lastIndex+1:]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		app.jsonOuput(w, "MAD", http.StatusBadRequest, nil)
	}

	user := struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}{
		Id:   int64(id),
		Name: "user",
	}

	app.jsonOuput(w, user, http.StatusOK, nil)

}
