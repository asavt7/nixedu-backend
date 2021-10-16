package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
	PostsStorage
	CommentsStorage
	UserStorage
}

func NewPostgresStorage(db *sql.DB) *Storage {

	return &Storage{
		PostsStorage:    &PostgresPostsStorage{db: db},
		CommentsStorage: &PostgresCommentsStorage{db: db},
		UserStorage:     &PostgresUserStorage{db: db},
	}
}
