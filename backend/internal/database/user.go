package database

import (
	"AUThConnect/internal/validator"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Password  password  `json:"-"`
	Role      string    `json:"role,omitempty"`
	Email     string    `json:"email"`
	Activated bool      `json:"activated"`
	Version   int       `json:"-"`
	CreatedAt time.Time `json:"-"`
}

type password struct {
	text *string
	hash []byte
}

func (p *password) Set(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	p.text = &password
	p.hash = hash

	return nil
}

func (p *password) Matches(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(password))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(validator.Matches(email, validator.EmailRX), "email", "must be a valid email address")
}

func ValidatePassword(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 characters long")
	v.Check(len(password) <= 72, "password", "must not be more than 72 characters long")
}

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.Name != "", "name", "must be provided")
	v.Check(len(user.Name) <= 50, "name", "must not be more than 50 characters long")

	ValidateEmail(v, user.Email)

	if user.Password.text == nil {
		panic("missing password hash for user")
	}
}

func ValidateLogin(v *validator.Validator, name, password string) {
	v.Check(name != "", "name", "must be provided")
	v.Check(len(name) <= 50, "name", "must not be more than 50 characters")

	ValidatePassword(v, password)
}

type UserModel struct {
	db *sql.DB
}

func (u UserModel) GetByEmail(email string) (*User, error) {
	query := `
    SELECT id, name, email, hashed_password, role, created_at
    FROM users
    WHERE email = $1
  `

	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := u.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password.hash,
		&user.Role,
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

func (u UserModel) Create(user *User) error {
	query := `
    INSERT INTO users (name, email, hashed_password, activated)
    VALUES ($1, $2, $3, $4)
    RETURNING id, created_at, version
  `

	args := []any{user.Name, user.Email, user.Password.hash, user.Activated}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var pqErr *pq.Error

	// TODO: Handle all possible errors
	err := u.db.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.CreatedAt, &user.Version)
	if err != nil {
		switch {
		case errors.As(err, &pqErr):
			if pqErr.Code == "23505" {
				return ErrUsernameOrEmailExists
			}
		default:
			return err
		}
	}

	return nil
}

func (u UserModel) Update(user *User) error {
	query := `
    UPDATE users
    SET name = $1, email = $2, hashed_password = $3, activated = $4, version = version + 1
    WHERE id = $5 AND version = $6
    RETURNING version
  `

	args := []any{
		user.Name,
		user.Email,
		user.Password.hash,
		user.Activated,
		user.ID,
		user.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var pqErr *pq.Error

	err := u.db.QueryRowContext(ctx, query, args...).Scan(&user.Version)
	if err != nil {
		switch {
		case errors.As(err, &pqErr):
			if pqErr.Code == "23505" {
				return ErrUsernameOrEmailExists
			}
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}
