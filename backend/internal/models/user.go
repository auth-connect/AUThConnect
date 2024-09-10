package models

import "fmt"

type User struct {
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
