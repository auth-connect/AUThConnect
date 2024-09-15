package models

import (
	"fmt"
	"time"
)

type User struct {
	ID             int64     `gorm:"primaryKey"`
	Username       string    `gorm:"type:varchar(255);uniqueIndex:idx_username_email"` // Index on username
	HashedPassword string    `gorm:"type:varchar(255);not null"`
	FullName       string    `gorm:"type:varchar(255);not null"`
	Role           string    `gorm:"type:varchar(255);not null"`
	Email          string    `gorm:"type:varchar(255);unique;not null"` // Unique constraint on email
	CreatedAt      time.Time `gorm:"type:timestamptz;not null;default:CURRENT_TIMESTAMP"`
}

type InputUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
	Email    string `json:"email"`
}

type ReturnUser struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
	Email    string `json:"email"`
}

// Formats the ReturnUser instance similar to the one being returned by the handler
func (ru *ReturnUser) String() string {
	return fmt.Sprintf(`{"id":%d,"username":"%s","full_name":"%s","role":"%s","email":"%s"}`,
		ru.Id,
		ru.Username,
		ru.FullName,
		ru.Role,
		ru.Email,
	)
}
