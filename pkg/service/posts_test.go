package service_test

import (
	"errors"
	mockstorage "github.com/asavt7/nixedu/backend/mocks/pkg/storage"
	"github.com/asavt7/nixedu/backend/pkg/model"
	"github.com/asavt7/nixedu/backend/pkg/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func initMocks(t *testing.T) (service.PostService, *mockstorage.MockPostsStorage, func()) {
	controller := gomock.NewController(t)
	finishFunc := controller.Finish

	postsStorage := mockstorage.NewMockPostsStorage(controller)
	postsService := service.NewPostServiceImpl(postsStorage)

	return postsService, postsStorage, finishFunc
}

func TestPostServiceImpl_GetAll(t *testing.T) {
	postsService, postsStorage, finishFunc := initMocks(t)
	defer finishFunc()

	posts := []model.Post{{
		UserID: 1,
		ID:     1,
		Title:  "title",
		Body:   "body",
	}}

	t.Run("ok", func(t *testing.T) {
		postsStorage.EXPECT().GetAll().Return(posts, nil)
		actual, err := postsService.GetAll()
		if err != nil {
			t.Errorf("err should be nil")
		}
		assert.Equal(t, posts, actual)
	})

	t.Run("ok", func(t *testing.T) {
		postsStorage.EXPECT().GetAll().Return(nil, errors.New("err"))
		_, err := postsService.GetAll()
		if err == nil {
			t.Errorf("err should not be nil")
		}
	})
}

func TestPostServiceImpl_GetAllByUserID(t *testing.T) {
	postsService, postsStorage, finishFunc := initMocks(t)
	defer finishFunc()
	posts := []model.Post{{
		UserID: 1,
		ID:     1,
		Title:  "title",
		Body:   "body",
	}}
	t.Run("ok", func(t *testing.T) {
		postsStorage.EXPECT().GetAllByUserID(1).Return(posts, nil)
		actual, err := postsService.GetAllByUserID(1)
		if err != nil {
			t.Errorf("err should be nil")
		}
		assert.Equal(t, posts, actual)
	})
	t.Run("err", func(t *testing.T) {
		postsStorage.EXPECT().GetAllByUserID(1).Return(nil, errors.New("err"))
		_, err := postsService.GetAllByUserID(1)
		if err == nil {
			t.Errorf("err should not be nil")
		}
	})
}

func TestPostServiceImpl_GetByUserIdAndPostId(t *testing.T) {
	postsService, postsStorage, finishFunc := initMocks(t)
	defer finishFunc()
	post := model.Post{
		UserID: 1,
		ID:     1,
		Title:  "title",
		Body:   "body",
	}
	t.Run("ok", func(t *testing.T) {
		postsStorage.EXPECT().GetByUserIDAndID(1, 1).Return(post, nil)
		actual, err := postsService.GetByUserIDAndPostID(1, 1)
		if err != nil {
			t.Errorf("err should be nil")
		}
		assert.Equal(t, post, actual)
	})
	t.Run("err", func(t *testing.T) {
		postsStorage.EXPECT().GetByUserIDAndID(1, 1).Return(model.Post{}, errors.New("err"))
		_, err := postsService.GetByUserIDAndPostID(1, 1)
		if err == nil {
			t.Errorf("err should not be nil")
		}
	})
}

func TestPostServiceImpl_DeletePost(t *testing.T) {
	postsService, postsStorage, finishFunc := initMocks(t)
	defer finishFunc()

	t.Run("ok", func(t *testing.T) {
		postsStorage.EXPECT().DeleteByUserIDAndID(1, 1).Return(nil)
		err := postsService.DeletePost(1, 1)
		if err != nil {
			t.Errorf("err should be nil")
		}
	})
	t.Run("err", func(t *testing.T) {
		postsStorage.EXPECT().DeleteByUserIDAndID(1, 1).Return(errors.New("cannot delete"))
		err := postsService.DeletePost(1, 1)
		if err == nil {
			t.Errorf("err should not be nil")
		}
	})
}

func TestPostServiceImpl_Save(t *testing.T) {
	postsService, postsStorage, finishFunc := initMocks(t)
	defer finishFunc()
	post := model.Post{
		UserID: 1,
		ID:     1,
		Title:  "title",
		Body:   "body",
	}
	t.Run("ok", func(t *testing.T) {
		postsStorage.EXPECT().Save(post).Return(post, nil)
		actual, err := postsService.Save(post)
		if err != nil {
			t.Errorf("err should be nil")
		}
		assert.Equal(t, post, actual)
	})
	t.Run("err", func(t *testing.T) {
		postsStorage.EXPECT().Save(post).Return(post, errors.New("cannot save"))
		_, err := postsService.Save(post)
		if err == nil {
			t.Errorf("err should not be nil")
		}
	})
}

func TestPostServiceImpl_Update(t *testing.T) {
	title := "new title"
	body := "new body"

	postsService, postsStorage, finishFunc := initMocks(t)
	defer finishFunc()

	update := model.UpdatePost{
		Title: &title,
		Body:  &body,
	}
	post := model.Post{
		UserID: 1,
		ID:     1,
		Title:  title,
		Body:   body,
	}
	t.Run("ok", func(t *testing.T) {
		postsStorage.EXPECT().Update(1, 1, update).Return(post, nil)
		actual, err := postsService.Update(1, 1, update)
		if err != nil {
			t.Errorf("err should be nil")
		}
		assert.Equal(t, post, actual)
	})
	t.Run("err", func(t *testing.T) {
		postsStorage.EXPECT().Update(1, 1, update).Return(post, errors.New("err updating "))
		_, err := postsService.Update(1, 1, update)
		if err == nil {
			t.Errorf("err should not be nil")
		}
	})
}
