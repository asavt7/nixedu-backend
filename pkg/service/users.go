package service

import (
	"github.com/asavt7/nixedu/backend/pkg/model"
	"github.com/asavt7/nixedu/backend/pkg/storage"
)

type UserServiceImpl struct {
	repo storage.UserStorage
}

func (u *UserServiceImpl) CreateUser(user model.User, password string) (model.User, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return model.User{}, err
	}
	user.PasswordHash = hashedPassword
	return u.repo.Create(user)
}

func (u *UserServiceImpl) GetUserByID(id int) (model.User, error) {
	return u.repo.GetByID(id)
}

func (u UserServiceImpl) GetUserByEmail(email string) (model.User, error) {
	return u.repo.FindByEmail(email)
}
