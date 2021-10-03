package service

import "github.com/asavt7/nixEducation/pkg/model"

type PostService interface {
	GetAll() ([]model.Post, error)
}

type Service interface {
	PostService
}
