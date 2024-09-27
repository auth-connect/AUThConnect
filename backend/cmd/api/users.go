package main

import (
	"AUThConnect/internal/models"
	"errors"
	"fmt"
	"net/http"
)

func (app *application) getUser(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}

	user, err := app.models.User.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	u := struct {
		ID       int64  `json:"id"`
		UserName string `json:"user_name"`
		FullName string `json:"full_name"`
		Role     string `json:"role"`
		Email    string `json:"email"`
	}{
		user.ID,
		user.UserName,
		user.FullName,
		user.Role,
		user.Email,
	}

	err = app.jsonWrite(w, envelope{"user": u}, http.StatusOK, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	var input models.InputUser

	err := app.jsonRead(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &models.InputUser{
		UserName: input.UserName,
		FullName: input.FullName,
		Password: input.Password,
		Email:    input.Email,
	}

	// TODO: Validate the payload

	id, err := app.models.User.Create(user)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrUsernameOrEmailExists):
			app.badRequestResponse(w, r, err)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/users/%d", id))

	err = app.jsonWrite(w, envelope{"user": user}, http.StatusCreated, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateUser(w http.ResponseWriter, r *http.Request) {
	var input models.InputUser

	id, err := app.readIDParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}

	err = app.jsonRead(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &models.InputUser{
		UserName: input.UserName,
		FullName: input.FullName,
		Password: input.Password,
		Email:    input.Email,
	}

	// TODO: Validate the payload

	err = app.models.User.Update(id, user)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrRecordNotFound):
			app.badRequestResponse(w, r, err)
		case errors.Is(err, models.ErrPasswordMismatch):
			app.badRequestResponse(w, r, err)
		case errors.Is(err, models.ErrUsernameOrEmailExists):
			app.badRequestResponse(w, r, err)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/users/%d", id))

	err = app.jsonWrite(w, envelope{"user": user}, http.StatusCreated, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
