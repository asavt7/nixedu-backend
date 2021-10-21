package service_test

import (
	"errors"
	mockstorage "github.com/asavt7/nixEducation/mocks/pkg/storage"
	"github.com/asavt7/nixEducation/pkg/model"
	"github.com/asavt7/nixEducation/pkg/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func initCommentsMocks(t *testing.T) (service.CommentService, *mockstorage.MockCommentsStorage, func()) {
	controller := gomock.NewController(t)
	finishFunc := controller.Finish

	postsStorage := mockstorage.NewMockCommentsStorage(controller)
	postsService := service.NewCommentServiceImpl(postsStorage)

	return postsService, postsStorage, finishFunc
}

func TestCommentServiceImpl_GetAllByPostID(t *testing.T) {
	s, mock, finishFunc := initCommentsMocks(t)
	defer finishFunc()

	comments := []model.Comment{{
		PostID: 1,
		Id:     1,
		UserID: 1,
		Body:   "qwe",
	}}

	t.Run("ok", func(t *testing.T) {
		mock.EXPECT().GetAllByPostID(1).Return(comments, nil)

		actual, err := s.GetAllByPostID(1)
		if err != nil {
			t.Errorf("err should be nil")
		}
		assert.Equal(t, comments, actual)
	})
	t.Run("err", func(t *testing.T) {
		mock.EXPECT().GetAllByPostID(1).Return(comments, model.PostNotFoundErr{Id: 1})
		_, err := s.GetAllByPostID(1)
		if err == nil {
			t.Errorf("err should not be nil")
		}
	})
}

func TestCommentServiceImpl_Delete(t *testing.T) {
	s, mock, finishFunc := initCommentsMocks(t)
	defer finishFunc()

	t.Run("ok", func(t *testing.T) {
		mock.EXPECT().DeleteByUserIDAndID(1, 1).Return(nil)

		err := s.Delete(1, 1)
		if err != nil {
			t.Errorf("err should be nil")
		}
	})
	t.Run("err", func(t *testing.T) {
		mock.EXPECT().DeleteByUserIDAndID(1, 1).Return(model.CommentNotFoundErr{Id: 1})
		err := s.Delete(1, 1)
		if err == nil {
			t.Errorf("err should not be nil")
		}
	})
}

func TestCommentServiceImpl_Save(t *testing.T) {
	s, mock, finishFunc := initCommentsMocks(t)
	defer finishFunc()

	comment := model.Comment{
		PostID: 1,
		Id:     1,
		UserID: 1,
		Body:   "qwe",
	}

	t.Run("ok", func(t *testing.T) {
		mock.EXPECT().Save(comment).Return(comment, nil)

		actual, err := s.Save(comment)
		if err != nil {
			t.Errorf("err should be nil")
		}
		assert.Equal(t, comment, actual)
	})
	t.Run("err", func(t *testing.T) {
		mock.EXPECT().Save(comment).Return(comment, model.PostNotFoundErr{Id: 1})
		_, err := s.Save(comment)
		if err == nil {
			t.Errorf("err should not be nil")
		}
	})
}

func TestCommentServiceImpl_Update(t *testing.T) {
	s, mock, finishFunc := initCommentsMocks(t)
	defer finishFunc()

	body := "nil"
	update := model.UpdateComment{Body: &body}

	comment := model.Comment{
		PostID: 1,
		Id:     1,
		UserID: 1,
		Body:   body,
	}

	t.Run("ok", func(t *testing.T) {
		mock.EXPECT().Update(1, 1, update).Return(comment, nil)

		actual, err := s.Update(1, 1, update)
		if err != nil {
			t.Errorf("err should be nil")
		}
		assert.Equal(t, comment, actual)
	})
	t.Run("err", func(t *testing.T) {
		mock.EXPECT().Update(1, 1, update).Return(comment, errors.New("cannot update"))
		_, err := s.Update(1, 1, update)
		if err == nil {
			t.Errorf("err should not be nil")
		}
	})
}
