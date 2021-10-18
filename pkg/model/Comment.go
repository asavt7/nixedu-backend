package model

type Comment struct {
	PostId int    `json:"postId" xml:"postId"`
	Id     int    `json:"id" xml:"id"`
	Name   string `json:"name" xml:"name"`
	Email  string `json:"email" xml:"email"`
	Body   string `json:"body" xml:"body"`
}

type UpdateComment struct {
	Body *string `json:"body" xml:"body"`
}
