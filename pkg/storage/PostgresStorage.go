package storage

import (
	"database/sql"
	"github.com/asavt7/nixEducation/pkg/model"
	_ "github.com/lib/pq"
)

type PostgresCommentsStorage struct {
	db *sql.DB
}

func (p *PostgresCommentsStorage) SaveAll(comments []model.Comment) ([]model.Comment, error) {
	panic("implement me")
}

func (p *PostgresCommentsStorage) Save(c model.Comment) (model.Comment, error) {

	sqlStatement :=
		`INSERT INTO nix.comments (id, postId, name, email, Body) VALUES ($1, $2, $3, $4, $5) returning *`

	res := p.db.QueryRow(sqlStatement, c.Id, c.PostId, c.Name, c.Email, c.Body)
	err := res.Scan(&c.Id, &c.PostId, &c.Name, &c.Email, &c.Body)
	if err != nil {
		return model.Comment{}, err
	}
	return c, nil
}

type PostgresPostsStorage struct {
	db *sql.DB
}

func (p *PostgresPostsStorage) SaveAll(posts []model.Post) ([]model.Post, error) {
	panic("implement me")
}

func (p *PostgresPostsStorage) Save(post model.Post) (model.Post, error) {
	sqlStatement :=
		`INSERT INTO nix.posts (id, UserId, Title, Body) VALUES ($1, $2, $3, $4) returning *`

	res := p.db.QueryRow(sqlStatement, post.Id, post.UserId, post.Title, post.Body)
	err := res.Scan(&post.Id, &post.UserId, &post.Title, &post.Body)
	if err != nil {
		return model.Post{}, err
	}
	return post, nil

}

type PostgresStorage struct {
	db *sql.DB
	PostsStorage
	CommentsStorage
}

func NewPostgresStorage(db *sql.DB) *PostgresStorage {

	return &PostgresStorage{
		db:              db,
		PostsStorage:    &PostgresPostsStorage{db: db},
		CommentsStorage: &PostgresCommentsStorage{db: db},
	}
}
