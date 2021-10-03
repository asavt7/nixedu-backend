package storage

import "github.com/asavt7/nixEducation/pkg/model"

type PostsStorage interface {
	SaveAll(posts []model.Post) ([]model.Post, error)
	Save(p model.Post) (model.Post, error)
}

type CommentsStorage interface {
	SaveAll(posts []model.Comment) ([]model.Comment, error)
	Save(c model.Comment) (model.Comment, error)
}

type Storage struct {
	PostsStorage
	CommentsStorage
}
