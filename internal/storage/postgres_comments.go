package storage

import (
	"errors"
	"fmt"
	"github.com/asavt7/nixedu/backend/internal/model"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type PostgresCommentsStorage struct {
	db *sqlx.DB
}

func (p *PostgresCommentsStorage) GetAllByUserID(userID int) ([]model.Comment, error) {
	var comments []model.Comment
	query := fmt.Sprintf("SELECT * FROM %s WHERE userid=$1", commentsTable)
	err := p.db.Select(&comments, query, userID)
	if err != nil {
		log.Error(err.Error())
	}
	return comments, err
}

func (p *PostgresCommentsStorage) GetAllByPostID(postID int) ([]model.Comment, error) {
	var comments []model.Comment
	query := fmt.Sprintf("SELECT * FROM %s WHERE postid=$1", commentsTable)
	err := p.db.Select(&comments, query, postID)
	if err != nil {
		log.Error(err.Error())
	}
	return comments, err
}

func (p *PostgresCommentsStorage) GetByCommentID(commentID int) (model.Comment, error) {
	var comment *model.Comment
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", commentsTable)
	err := p.db.Get(&comment, query, commentID)
	if err != nil {
		log.Error(err.Error())
	}
	if comment == nil {
		return *comment, model.CommentNotFoundErr{ID: commentID}
	}
	return *comment, err
}

func (p *PostgresCommentsStorage) Update(userID, commentID int, comment model.UpdateComment) (model.Comment, error) {

	var result model.Comment
	if comment.Body == nil {
		return result, errors.New("empty fields to update")
	}

	argNum := 0
	updateArgs := make([]string, 0)
	updateVals := make([]interface{}, 0)

	if comment.Body != nil {
		updateArgs = append(updateArgs, "body")
		updateVals = append(updateVals, *comment.Body)
		argNum++
	}

	updateVals = append(updateVals, userID, commentID)

	setExpression := convertToSetStrs(updateArgs)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE userid=$%d AND id=$%d RETURNING *", postsTable, setExpression, argNum+1, argNum+2)

	err := p.db.Get(&result, query, updateVals...)
	if err != nil {
		log.Error(err.Error())
	}
	return result, err

}

func (p *PostgresCommentsStorage) DeleteByUserIDAndID(userID, commentID int) error {
	var id *int
	query := fmt.Sprintf("DELETE FROM %s WHERE userid=$1 AND id=$2 RETURNING id", commentsTable)
	err := p.db.Get(&id, query, userID, commentID)
	if err != nil {
		log.Error(err.Error())
	}
	if id == nil {
		return model.CommentNotFoundErr{ID: commentID}
	}
	return err
}

func (p *PostgresCommentsStorage) Save(c model.Comment) (model.Comment, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (userid, postid, body) VALUES ($1,$2,$3) RETURNING id", commentsTable)
	err := p.db.Get(&id, query, c.UserID, c.PostID, c.Body)
	c.ID = id
	if err != nil {
		log.Error(err.Error())
	}
	return c, err
}
