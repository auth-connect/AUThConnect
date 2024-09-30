package main

import (
	"AUThConnect/internal/database"
	"AUThConnect/internal/validator"
	"errors"
	"net/http"
)

func (app *application) registerUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.jsonRead(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &database.User{
		Name:      input.Name,
		Email:     input.Email,
		Activated: false,
	}

	err = user.Password.Set(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	v := validator.New()

	if database.ValidateUser(v, user); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Users.Create(user)
	if err != nil {
		switch {
		case errors.Is(err, database.ErrUsernameOrEmailExists):
			v.AddError("email", "a user with this username/email address already exists")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.jsonWrite(w, envelope{"user": user}, http.StatusCreated, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// func (app *application) activateUser(w http.ResponseWriter, r *http.Request) {
//   var input struct {
//
//   }
// }
