package models

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
