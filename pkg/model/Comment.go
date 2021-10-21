package model

// Comment - user comment
type Comment struct {
	PostId int    `json:"postId" xml:"postId" validate:"required"`
	Id     int    `json:"id" xml:"id"`
	UserID int    `json:"userId" xml:"userId" validate:"required"`
	Body   string `json:"body" xml:"body" validate:"required"`
}

// UpdateComment - info for updating user comment
type UpdateComment struct {
	Body *string `json:"body" xml:"body" validate:"required"`
}
