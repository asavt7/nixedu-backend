package storage

import (
	"github.com/jmoiron/sqlx"
	// import postgres drivers
	_ "github.com/lib/pq"
)

const (
	usersTable = "nix.users"
	postsTable = "nix.posts"
)

// PostgresStorage - postgres storage implementation
type PostgresStorage struct {
	db *sqlx.DB
	PostsStorage
	CommentsStorage
	UserStorage
}

// NewPostgresStorage constructs PostgresStorage instance
func NewPostgresStorage(db *sqlx.DB) *Storage {

	return &Storage{
		PostsStorage:    &PostgresPostsStorage{db: db},
		CommentsStorage: &PostgresCommentsStorage{db: db},
		UserStorage:     &PostgresUserStorage{db: db},
	}
}
