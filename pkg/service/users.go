package service

import (
	"github.com/asavt7/nixEducation/pkg/model"
	"github.com/asavt7/nixEducation/pkg/storage"
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

func (u *UserServiceImpl) GetUserById(id int) (model.User, error) {
	return u.repo.GetByID(id)
}

func (u UserServiceImpl) GetUserByEmail(email string) (model.User, error) {
	return u.repo.FindByEmail(email)
}
