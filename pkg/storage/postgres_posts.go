package storage

import (
	"errors"
	"fmt"
	"github.com/asavt7/nixEducation/pkg/model"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"strings"
)

type PostgresPostsStorage struct {
	db *sqlx.DB
}

func (p *PostgresPostsStorage) GetAll() ([]model.Post, error) {
	var posts []model.Post
	query := fmt.Sprintf("SELECT * FROM %s", postsTable)
	err := p.db.Select(&posts, query)
	if err != nil {
		log.Error(err.Error())
	}
	return posts, err
}

func (p *PostgresPostsStorage) Save(post model.Post) (model.Post, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (userid, title, body) VALUES ($1,$2,$3) RETURNING id", postsTable)
	err := p.db.Get(&id, query, post.UserID, post.Title, post.Body)
	post.ID = id
	if err != nil {
		log.Error(err.Error())
	}
	return post, err
}

func (p *PostgresPostsStorage) GetAllByUserID(userID int) ([]model.Post, error) {
	var posts []model.Post
	query := fmt.Sprintf("SELECT * FROM %s WHERE userid=$1", postsTable)
	err := p.db.Select(&posts, query, userID)
	if err != nil {
		log.Error(err.Error())
	}
	return posts, err
}

func (p *PostgresPostsStorage) GetByUserIDAndID(userID, postID int) (model.Post, error) {
	var post model.Post
	query := fmt.Sprintf("SELECT * FROM %s WHERE userid=$1 AND id=$2", postsTable)
	err := p.db.Get(&post, query, userID, postID)
	if err != nil {
		log.Error(err.Error())
	}
	return post, err
}

func (p *PostgresPostsStorage) Update(userID, postID int, post model.UpdatePost) (model.Post, error) {
	var resultPost model.Post
	if post.Body == nil && post.Title == nil {
		return resultPost, errors.New("empty fields to update")
	}

	argNum := 0
	updateArgs := make([]string, 0)
	updateVals := make([]interface{}, 0)

	if post.Title != nil {
		updateArgs = append(updateArgs, "title")
		updateVals = append(updateVals, *post.Title)
		argNum++
	}
	if post.Body != nil {
		updateArgs = append(updateArgs, "body")
		updateVals = append(updateVals, *post.Body)
		argNum++
	}

	updateVals = append(updateVals, userID, postID)

	setExpression := convertToSetStrs(updateArgs)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE userid=$%d AND id=$%d", postsTable, setExpression, argNum+1, argNum+2)

	err := p.db.Get(&resultPost, query, updateVals...)
	if err != nil {
		log.Error(err.Error())
	}
	return resultPost, err
}

func convertToSetStrs(args []string) string {
	strs := make([]string, 0)
	for i, v := range args {
		strs = append(strs, fmt.Sprintf("%s=$%d", v, i+1))
	}
	return strings.Join(strs, ", ")
}

func (p *PostgresPostsStorage) DeleteByUserIDAndID(userID, postID int) error {
	var id int
	query := fmt.Sprintf("DELETE FROM %s WHERE userid=$1 AND id=$2 RETURNING id", postsTable)
	err := p.db.Get(&id, query, userID, postID)
	if err != nil {
		log.Error(err.Error())
	}
	if &id == nil {
		return model.PostNotFoundErr{Id: postID}
	}
	return err
}
