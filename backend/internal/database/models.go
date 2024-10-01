package database

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound        = errors.New("record not found")
	ErrPasswordMismatch      = errors.New("passwords do not match")
	ErrUsernameOrEmailExists = errors.New("a user with the same username or email already exists")
	ErrEditConflict          = errors.New("edit conflict")
)

type Models struct {
	Users UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Users: UserModel{db: db},
	}
}
