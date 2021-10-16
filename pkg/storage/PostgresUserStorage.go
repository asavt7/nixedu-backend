package storage

import (
	"database/sql"
	"github.com/asavt7/nixEducation/pkg/model"
)

type PostgresUserStorage struct {
	db *sql.DB
}

func (p PostgresUserStorage) GetByUsernameAndPasswordHash(username, passwordHash string) (model.User, error) {
	panic("implement me")
}

func (p PostgresUserStorage) Create(user model.User) (model.User, error) {
	panic("implement me")
}

func (p PostgresUserStorage) GetById(userId int) (model.User, error) {
	panic("implement me")
}

func (p PostgresUserStorage) GetByUsername(username string) (model.User, error) {
	panic("implement me")
}
