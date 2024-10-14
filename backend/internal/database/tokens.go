package database

import (
	"AUThConnect/internal/validator"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base32"
	"time"
)

const (
	ScopeActivation     = "activation"
	ScopeAuthentication = "authentication"
)

type Token struct {
	Text   string    `json:"text"`
	Hash   []byte    `json:"-"`
	UserID int64     `json:"-"`
	Expiry time.Time `json:"expiry"`
	Scope  string    `json:"-"`
}

func generateToken(userID int64, ttl time.Duration, scope string) (*Token, error) {
	token := &Token{
		UserID: userID,
		Expiry: time.Now().Add(ttl),
		Scope:  scope,
	}

	bytes := make([]byte, 16)

	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}

	token.Text = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(bytes)

	hash := sha256.Sum256([]byte(token.Text))
	token.Hash = hash[:]

	return token, nil
}

func ValidateTokenText(v *validator.Validator, tokenText string) {
	v.Check(tokenText != "", "token", "must be provided")
	v.Check(len(tokenText) == 26, "token", "must be 26 bytes long")
}

type TokenModel struct {
	db *sql.DB
}

func (t TokenModel) New(userID int64, ttl time.Duration, scope string) (*Token, error) {
	token, err := generateToken(userID, ttl, scope)
	if err != nil {
		return nil, err
	}

	err = t.Insert(token)
	return token, err
}

func (t TokenModel) Insert(token *Token) error {
	query := `
    INSERT INTO tokens (hash, user_id, expiry, scope)
    VALUES ($1, $2, $3, $4)
  `

	args := []any{token.Hash, token.UserID, token.Expiry, token.Scope}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := t.db.ExecContext(ctx, query, args...)
	return err
}

func (t TokenModel) DeleteAllForUser(scope string, userID int64) error {
	query := `
    DELETE FROM tokens
    WHERE scope = $1 AND user_id = $2;
  `

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := t.db.ExecContext(ctx, query, scope, userID)
	return err
}
