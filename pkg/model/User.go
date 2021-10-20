package model

// User model
type User struct {
	Id           int    `json:"id" xml:"id"`
	Username     string `json:"username" xml:"username" binding:"required" validate:"required"`
	Email        string `json:"email" xml:"email" binding:"required" validate:"required,email"`
	PasswordHash string `json:"-" xml:"-" db:"password_hash"`
}
