package model

// Post - user post
type Post struct {
	UserID int    `json:"userId" xml:"userId"`
	ID     int    `json:"id" xml:"id"`
	Title  string `json:"title" xml:"title" validate:"required"`
	Body   string `json:"body" xml:"body" validate:"required"`
}

// UpdatePost - udate post info
type UpdatePost struct {
	Title *string `json:"title" xml:"title" validate:"required"`
	Body  *string `json:"body" xml:"body" validate:"required"`
}
