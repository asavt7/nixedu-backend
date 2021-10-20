package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"

	"github.com/asavt7/nixEducation/pkg/model"
)

type PostgresUserStorage struct {
	db *sqlx.DB
}

func (p *PostgresUserStorage) FindByUsername(username string) (model.User, error) {
	var user model.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1", usersTable)
	err := p.db.Get(&user, query, username)
	if err == sql.ErrNoRows {
		return user, model.UserNotFoundErr{}
	}
	if err != nil {
		log.Error(err.Error())
	}
	return user, err

}

func (p *PostgresUserStorage) FindByEmail(email string) (model.User, error) {
	var user model.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE email=$1", usersTable)
	err := p.db.Get(&user, query, email)
	if err == sql.ErrNoRows {
		return user, model.UserNotFoundErr{}
	}
	if err != nil {
		log.Error(err.Error())
	}
	return user, err
}

func (p *PostgresUserStorage) GetByUsernameAndPasswordHash(username, passwordHash string) (model.User, error) {
	var user model.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1 and password_hash=$2", usersTable)
	err := p.db.Get(&user, query, username, passwordHash)
	if err == sql.ErrNoRows {
		return user, model.UserNotFoundErr{}
	}
	if err != nil {
		log.Error(err.Error())
		errors.New(fmt.Sprintf("cannot find user by username=%s and password", username))
	}
	return user, err
}

func (p *PostgresUserStorage) Create(user model.User) (model.User, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (username,email,password_hash) values( $1,$2, $3) RETURNING id", usersTable)
	err := p.db.Get(&id, query, user.Username, user.Email, user.PasswordHash)
	if err != nil {
		log.Error(err.Error())
		return user, errors.New(fmt.Sprintf("cannot create user username=%s email=%s", user.Username, user.Email))
	}
	user.Id = id
	return user, nil

}

func (p *PostgresUserStorage) GetById(userId int) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", usersTable)
	err := p.db.Get(&user, query, userId)
	if err == sql.ErrNoRows {
		return user, model.UserNotFoundErr{}
	}
	if err != nil {
		log.Error(err.Error())
	}
	return user, err
}