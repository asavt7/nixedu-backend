package storage

import "github.com/asavt7/nixEducation/pkg/model"

type PostsStorage interface {
	SaveAll(posts []model.Post) ([]model.Post, error)
}
type Storage struct {
	PostsStorage
}
