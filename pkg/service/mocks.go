package service

import "github.com/asavt7/nixEducation/pkg/model"

type MockUserService struct {
	CreateUserFunc func(user model.User, password string) (model.User, error)
	GetUserFunc    func(id int) (model.User, error)
}

func (m *MockUserService) CreateUser(user model.User, password string) (model.User, error) {
	return m.CreateUserFunc(user, password)
}

func (m *MockUserService) GetUser(id int) (model.User, error) {
	return m.GetUserFunc(id)
}
