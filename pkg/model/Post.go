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

type Album struct {
	UserID int    `json:"userId" xml:"userId"`
	Id     int    `json:"id" xml:"id"`
	Title  string `json:"title" xml:"title"`
}

type Photo struct {
	AlbumId      int    `json:"albumId" xml:"albumId"`
	Id           int    `json:"id" xml:"id"`
	Title        string `json:"title" xml:"title"`
	Url          string `json:"url" xml:"url"`
	ThumbnailUrl string `json:"thumbnailUrl" xml:"thumbnailUrl"`
}

type Todo struct {
	UserId    int    `json:"userId" xml:"userId"`
	Id        int    `json:"id" xml:"id"`
	Title     string `json:"title" xml:"title"`
	Completed bool   `json:"completed" xml:"completed"`
}
