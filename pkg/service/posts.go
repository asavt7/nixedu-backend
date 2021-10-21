package service

import (
	"github.com/asavt7/nixEducation/pkg/model"
	"github.com/asavt7/nixEducation/pkg/storage"
)

// PostServiceImpl implements service.PostService
type PostServiceImpl struct {
	repo storage.PostsStorage
}

// NewPostServiceImpl constructs PostServiceImpl instance
func NewPostServiceImpl(repo storage.PostsStorage) *PostServiceImpl {
	return &PostServiceImpl{repo: repo}
}

func (p *PostServiceImpl) GetAll() ([]model.Post, error) {
	panic("implement me")
}

func (p *PostServiceImpl) GetAllByUserID(userID int) ([]model.Post, error) {
	panic("implement me")
}

func (p *PostServiceImpl) GetByUserIDAndPostID(userID, postID int) (model.Post, error) {
	panic("implement me")
}

func (p *PostServiceImpl) Save(userID int, post model.Post) (model.Post, error) {
	panic("implement me")
}

func (p *PostServiceImpl) Update(userID, postID int, updatePost model.UpdatePost) (model.Post, error) {
	panic("implement me")
}

func (p *PostServiceImpl) DeletePost(userID, postID int) error {
	panic("implement me")
}
