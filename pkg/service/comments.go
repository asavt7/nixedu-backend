package service

import (
	"github.com/asavt7/nixEducation/pkg/model"
	"github.com/asavt7/nixEducation/pkg/storage"
)

type CommentServiceImpl struct {
	repo storage.CommentsStorage
}

func NewCommentServiceImpl(repo storage.CommentsStorage) *CommentServiceImpl {
	return &CommentServiceImpl{repo: repo}
}

func (c *CommentServiceImpl) GetAllByPostID(postID int) ([]model.Comment, error) {
	panic("implement me")
}

func (c *CommentServiceImpl) Save(comment model.Comment) (model.Comment, error) {
	panic("implement me")
}

func (c *CommentServiceImpl) Update(currentUserID, commentID int, comment model.UpdateComment) (model.Comment, error) {
	panic("implement me")
}

func (c *CommentServiceImpl) Delete(currentUserID, commentID int) error {
	panic("implement me")
}
