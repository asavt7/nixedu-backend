package service

import "github.com/asavt7/nixEducation/pkg/model"

type PostService interface {
	GetAll() ([]model.Post, error)
	GetAllByUserId(userId int) ([]model.Post, error)
}

type CommentService interface {
	GetAllByPostId(id int) ([]model.Comment, error)
}

type Service interface {
	PostService
	CommentService
}
