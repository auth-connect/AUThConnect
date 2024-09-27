package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int64     `json:"id"`
	UserName       string    `json:"user_name"`
	FullName       string    `json:"full_name"`
	HashedPassword string    `json:"hashed_password"`
	Role           string    `json:"role,omitempty"`
	Email          string    `json:"email"`
	CreatedAt      time.Time `json:"-"`
}

type InputUser struct {
	UserName string `json:"user_name"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserModel struct {
	db *sql.DB
}

func (u UserModel) Get(id int64) (*User, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
    SELECT id, user_name, full_name, role, email, created_at
    FROM users
    WHERE id = $1
  `

	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := u.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.UserName,
		&user.FullName,
		&user.Role,
		&user.Email,
		&user.CreatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (u UserModel) Create(input *InputUser) (int64, error) {
	query := `
    INSERT INTO users (user_name, full_name, hashed_password, email)
    VALUES ($1, $2, $3, $4)
    RETURNING id, created_at
  `

	hashed_password, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		return 0, err
	}

	user := User{
		UserName:       input.UserName,
		FullName:       input.FullName,
		HashedPassword: string(hashed_password),
		Email:          input.Email,
	}

	args := []any{user.UserName, user.FullName, user.HashedPassword, user.Email}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// TODO: Handle all possible errors
	err = u.db.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return 0, ErrUsernameOrEmailExists
		}

		return 0, fmt.Errorf("error inserting user: %w", err)
	}

	return user.ID, nil
}

func (u UserModel) Update(id int64, input *InputUser) error {
	query := `
    SELECT hashed_password
    FROM users
    WHERE id = $1
  `

	var hashed_password string

	err := u.db.QueryRow(query, id).Scan(&hashed_password)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(input.Password))
	if err != nil {
		return ErrPasswordMismatch
	}

	query = `
    UPDATE users
    SET user_name = $1, full_name = $2, hashed_password = $3, email = $4
    WHERE id = $5
  `

	args := []any{input.UserName, input.FullName, hashed_password, input.Email, id}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := u.db.ExecContext(ctx, query, args...)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return ErrUsernameOrEmailExists
		}

		return fmt.Errorf("error updating user: %w", err)
	}

	rowsAffecter, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffecter == 0 {
		return ErrRecordNotFound
	}

	return nil
}
