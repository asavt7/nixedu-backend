package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	usersTable = "nix.users"
)

type PostgresStorage struct {
	db *sqlx.DB
	PostsStorage
	CommentsStorage
	UserStorage
}

func NewPostgresStorage(db *sqlx.DB) *Storage {

	return &Storage{
		PostsStorage:    &PostgresPostsStorage{db: db},
		CommentsStorage: &PostgresCommentsStorage{db: db},
		UserStorage:     &PostgresUserStorage{db: db},
	}
}
