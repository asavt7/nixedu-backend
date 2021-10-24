package model

// Comment - user comment
type Comment struct {
	PostID int    `json:"postId" xml:"postId" validate:"required"`
	ID     int    `json:"id" xml:"id"`
	UserID int    `json:"userId" xml:"userId" validate:"required"`
	Body   string `json:"body" xml:"body" validate:"required"`
}

// UpdateComment - info for updating user comment
type UpdateComment struct {
	Body *string `json:"body" xml:"body" validate:"required"`
}
