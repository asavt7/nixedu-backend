package storage

import (
	"github.com/asavt7/nixEducation/pkg/model"
	"github.com/jmoiron/sqlx"
)

type PostgresCommentsStorage struct {
	db *sqlx.DB
}

func (p *PostgresCommentsStorage) GetAllByUserID(userID int) ([]model.Comment, error) {
	panic("implement me")
}

func (p *PostgresCommentsStorage) GetAllByPostID(postID int) ([]model.Comment, error) {
	panic("implement me")
}

func (p *PostgresCommentsStorage) GetByCommentID(commentID int) (model.Comment, error) {
	panic("implement me")
}

func (p *PostgresCommentsStorage) Update(userID, commentID int, comment model.UpdateComment) (model.Comment, error) {
	panic("implement me")
}

func (p *PostgresCommentsStorage) DeleteByUserIDAndID(userID, commentID int) error {
	panic("implement me")
}

func (p *PostgresCommentsStorage) SaveAll(comments []model.Comment) ([]model.Comment, error) {
	panic("implement me")
}

func (p *PostgresCommentsStorage) Save(c model.Comment) (model.Comment, error) {
	/*
		sqlStatement :=
			`INSERT INTO nix.comments (id, postId, name, email, Body) VALUES ($1, $2, $3, $4, $5) returning *`

		res := p.db.QueryRow(sqlStatement, c.ID, c.PostID, c.Name, c.Email, c.Body)
		err := res.Scan(&c.ID, &c.PostID, &c.Name, &c.Email, &c.Body)
		if err != nil {
			return model.Comment{}, err
		}
		return c, nil */
	panic("not impl")
}
