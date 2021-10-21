package storage

import (
	"github.com/asavt7/nixEducation/pkg/model"
	"github.com/jmoiron/sqlx"
)

type PostgresPostsStorage struct {
	db *sqlx.DB
}

func (p *PostgresPostsStorage) GetAll() ([]model.Post, error) {
	panic("implement me")
}

func (p *PostgresPostsStorage) Save(userID int, post model.Post) (model.Post, error) {
	panic("implement me")
}

func (p *PostgresPostsStorage) GetAllByUserID(userID int) ([]model.Post, error) {
	panic("implement me")
}

func (p *PostgresPostsStorage) GetByUserIDAndID(userID, postID int) (model.Post, error) {
	panic("implement me")
}

func (p *PostgresPostsStorage) Update(userID, postID int, post model.UpdatePost) (model.Post, error) {
	panic("implement me")
}

func (p *PostgresPostsStorage) DeleteByUserIDAndID(userID, postID int) error {
	panic("implement me")
}

/*
func (p *PostgresPostsStorage) Save(post model.Post) (model.Post, error) {
	sqlStatement :=
		`INSERT INTO nix.posts (id, UserId, Title, Body) VALUES ($1, $2, $3, $4) returning *`

	res := p.db.QueryRow(sqlStatement, post.ID, post.UserId, post.Title, post.Body)
	err := res.Scan(&post.ID, &post.UserId, &post.Title, &post.Body)
	if err != nil {
		return model.Post{}, err
	}
	return post, nil

}
*/
