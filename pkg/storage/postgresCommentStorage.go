package storage

import (
	"database/sql"
	"github.com/asavt7/nixEducation/pkg/model"
)

type PostgresCommentsStorage struct {
	db *sql.DB
}

func (p *PostgresCommentsStorage) GetAllByUserId(userId int) ([]model.Comment, error) {
	panic("implement me")
}

func (p *PostgresCommentsStorage) GetAllByPostId(postId int) ([]model.Comment, error) {
	panic("implement me")
}

func (p *PostgresCommentsStorage) GetByCommentId(commentId int) (model.Comment, error) {
	panic("implement me")
}

func (p *PostgresCommentsStorage) Update(userId, postId, commentId int, comment model.UpdateComment) (model.Comment, error) {
	panic("implement me")
}

func (p *PostgresCommentsStorage) DeleteByUserIdAndId(userId, commentId int) error {
	panic("implement me")
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
