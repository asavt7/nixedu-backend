package service

import (
	"github.com/asavt7/nixedu/backend/internal/model"
	"github.com/asavt7/nixedu/backend/internal/storage"
)

type CommentServiceImpl struct {
	repo storage.CommentsStorage
}

func NewCommentServiceImpl(repo storage.CommentsStorage) *CommentServiceImpl {
	return &CommentServiceImpl{repo: repo}
}

func (c *CommentServiceImpl) GetAllByPostID(postID int) ([]model.Comment, error) {
	return c.repo.GetAllByPostID(postID)
}

func (c *CommentServiceImpl) Save(comment model.Comment) (model.Comment, error) {
	return c.repo.Save(comment)
}

func (c *CommentServiceImpl) Update(currentUserID, commentID int, comment model.UpdateComment) (model.Comment, error) {
	return c.repo.Update(currentUserID, commentID, comment)
}

func (c *CommentServiceImpl) Delete(currentUserID, commentID int) error {
	return c.repo.DeleteByUserIDAndID(currentUserID, commentID)
}
