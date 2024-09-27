package models

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound        = errors.New("record not found")
	ErrPasswordMismatch      = errors.New("passwords do not match")
	ErrUsernameOrEmailExists = errors.New("a user with the same username or email already exists")
)

type Models struct {
	User UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		User: UserModel{db: db},
	}
}
