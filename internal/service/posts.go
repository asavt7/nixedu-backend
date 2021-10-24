package service

import (
	"github.com/asavt7/nixedu/backend/internal/model"
	"github.com/asavt7/nixedu/backend/internal/storage"
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
	return p.repo.GetAll()
}

func (p *PostServiceImpl) GetAllByUserID(userID int) ([]model.Post, error) {
	return p.repo.GetAllByUserID(userID)
}

func (p *PostServiceImpl) GetByUserIDAndPostID(userID, postID int) (model.Post, error) {
	return p.repo.GetByUserIDAndID(userID, postID)
}

func (p *PostServiceImpl) Save(post model.Post) (model.Post, error) {
	return p.repo.Save(post)
}

func (p *PostServiceImpl) Update(userID, postID int, updatePost model.UpdatePost) (model.Post, error) {
	return p.repo.Update(userID, postID, updatePost)
}

func (p *PostServiceImpl) DeletePost(userID, postID int) error {
	return p.repo.DeleteByUserIDAndID(userID, postID)
}
