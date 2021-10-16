package storage

import (
	"database/sql"
	"github.com/asavt7/nixEducation/pkg/model"
)

type PostgresPostsStorage struct {
	db *sql.DB
}

func (p *PostgresPostsStorage) SaveAll(userId int, posts []model.Post) ([]model.Post, error) {
	panic("implement me")
}

func (p *PostgresPostsStorage) Save(userId int, post model.Post) (model.Post, error) {
	panic("implement me")
}

func (p *PostgresPostsStorage) GetAllByUserId(userId int) ([]model.Post, error) {
	panic("implement me")
}

func (p *PostgresPostsStorage) GetByUserIdAndId(userId, postId int) (model.Post, error) {
	panic("implement me")
}

func (p *PostgresPostsStorage) Update(userId, postId int, post model.UpdatePost) (model.Post, error) {
	panic("implement me")
}

func (p *PostgresPostsStorage) DeleteByUserIdAndId(userId, postId int) error {
	panic("implement me")
}

/*
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
*/
