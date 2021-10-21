package storage

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"

	"github.com/asavt7/nixEducation/pkg/model"
)

// PostgresUserStorage - postgres UserStorage implementation
type PostgresUserStorage struct {
	db *sqlx.DB
}

// FindByUsername - FindByUsername
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

// FindByEmail FindByEmail
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

// GetByUsernameAndPasswordHash - GetByUsernameAndPasswordHash
func (p *PostgresUserStorage) GetByUsernameAndPasswordHash(username, passwordHash string) (model.User, error) {
	var user model.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1 and password_hash=$2", usersTable)
	err := p.db.Get(&user, query, username, passwordHash)
	if err == sql.ErrNoRows {
		return user, model.UserNotFoundErr{}
	}
	if err != nil {
		log.Error(err.Error())
		return model.User{}, fmt.Errorf("cannot find user by username=%s and password", username)
	}
	return user, err
}

// Create - insert new user in DB
func (p *PostgresUserStorage) Create(user model.User) (model.User, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (username,email,password_hash) values( $1,$2, $3) RETURNING id", usersTable)
	err := p.db.Get(&id, query, user.Username, user.Email, user.PasswordHash)
	if err != nil {
		log.Error(err.Error())
		return user, fmt.Errorf("cannot create user username=%s email=%s", user.Username, user.Email)
	}
	user.ID = id
	return user, nil

}

// GetByID - GetByID
func (p *PostgresUserStorage) GetByID(userID int) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", usersTable)
	err := p.db.Get(&user, query, userID)
	if err == sql.ErrNoRows {
		return user, model.UserNotFoundErr{}
	}
	if err != nil {
		log.Error(err.Error())
	}
	return user, err
}
