package model

type User struct {
	Id           int    `json:"id" xml:"id"`
	Username     string `json:"username" xml:"username"`
	Email        string `json:"email" xml:"email"`
	PasswordHash string `json:"-" xml:"-"`
}
