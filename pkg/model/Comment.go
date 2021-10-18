package model

type Comment struct {
	PostId int    `json:"postId" xml:"postId" validate:"required"`
	Id     int    `json:"id" xml:"id"`
	UserId int    `json:"userId" xml:"userId" validate:"required"`
	Body   string `json:"body" xml:"body" validate:"required"`
}

type UpdateComment struct {
	Body *string `json:"body" xml:"body" validate:"required"`
}
