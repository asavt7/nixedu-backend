package server

import (
	"encoding/json"
	mock_service "github.com/asavt7/nixEducation/mocks/pkg/service"
	"github.com/asavt7/nixEducation/pkg/model"
	"github.com/asavt7/nixEducation/pkg/service"
	"github.com/golang/mock/gomock"
	"github.com/kinbiko/jsonassert"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestPostsHandler(t *testing.T) {

	userId := 1
	post := model.Post{UserId: userId, Id: 1, Title: "title", Body: "body"}
	posts := []model.Post{post}
	expectedPostsJson, err := json.Marshal(posts)
	if err != nil {
		t.Fatal(err)
	}
	expectedPostJson, err := json.Marshal(post)
	if err != nil {
		t.Fatal(err)
	}

	controller := gomock.NewController(t)
	defer controller.Finish()
	postsService := mock_service.NewMockPostService(controller)
	mockService := &service.Service{
		PostService: postsService,
	}
	handler := NewApiHandler(mockService)

	t.Run("create - ok", func(t *testing.T) {
		postsService.EXPECT().Save(userId, model.Post{
			Title: "title",
			Body:  "body",
		}).Return(post, nil)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"title":"title","body":"body"}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts")
		c.SetParamNames("userId")
		c.SetParamValues(strconv.Itoa(userId))

		if assert.NoError(t, handler.createPost(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			jsonassert.New(t).Assertf(rec.Body.String(), string(expectedPostJson))
		}
	})

	t.Run("create - bad request body", func(t *testing.T) {
		postsService.EXPECT().Save(userId, model.Post{}).Times(0).Return(post, nil)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts")
		c.SetParamNames("userId")
		c.SetParamValues(strconv.Itoa(userId))

		if assert.NoError(t, handler.createPost(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("get all user posts", func(t *testing.T) {
		postsService.EXPECT().GetAllByUserId(userId).Return(posts, nil)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts")
		c.SetParamNames("userId")
		c.SetParamValues(strconv.Itoa(userId))

		if assert.NoError(t, handler.getUserPosts(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			jsonassert.New(t).Assertf(rec.Body.String(), string(expectedPostsJson))
		}
	})

	t.Run("get all user posts - user not found", func(t *testing.T) {
		postsService.EXPECT().GetAllByUserId(userId).Return(nil, service.UserNotFoundErr{Id: userId})

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts")
		c.SetParamNames("userId")
		c.SetParamValues(strconv.Itoa(userId))

		if assert.NoError(t, handler.getUserPosts(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}
	})

	t.Run("get user post by id", func(t *testing.T) {
		postsService.EXPECT().GetByUserIdAndPostId(userId, post.Id).Return(post, nil)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts/:postId")
		c.SetParamNames("userId", "postId")
		c.SetParamValues(strconv.Itoa(userId), strconv.Itoa(post.Id))

		if assert.NoError(t, handler.getUserPostById(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			jsonassert.New(t).Assertf(rec.Body.String(), string(expectedPostJson))
		}
	})

	t.Run("get user post by id - user not found", func(t *testing.T) {
		postsService.EXPECT().GetByUserIdAndPostId(userId, post.Id).Return(model.Post{}, service.UserNotFoundErr{Id: userId})

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts/:postId")
		c.SetParamNames("userId", "postId")
		c.SetParamValues(strconv.Itoa(userId), strconv.Itoa(post.Id))

		if assert.NoError(t, handler.getUserPostById(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}
	})

	t.Run("get user post by id - post not found", func(t *testing.T) {
		postsService.EXPECT().GetByUserIdAndPostId(userId, post.Id).Return(model.Post{}, service.PostNotFoundErr{Id: post.Id})

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts/:postId")
		c.SetParamNames("userId", "postId")
		c.SetParamValues(strconv.Itoa(userId), strconv.Itoa(post.Id))

		if assert.NoError(t, handler.getUserPostById(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}
	})

	t.Run("update - ok", func(t *testing.T) {
		postsService.EXPECT().Update(userId, post.Id, model.UpdatePost{
			Title: &post.Title,
			Body:  &post.Body,
		}).Return(post, nil)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"title":"title","body":"body"}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts/:postId")
		c.SetParamNames("userId", "postId")
		c.SetParamValues(strconv.Itoa(userId), strconv.Itoa(post.Id))

		if assert.NoError(t, handler.updatePost(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			jsonassert.New(t).Assertf(rec.Body.String(), string(expectedPostJson))
		}
	})

	t.Run("update - bad request - empty body", func(t *testing.T) {
		postsService.EXPECT().Update(userId, post.Id, model.UpdatePost{
			Title: nil,
			Body:  nil,
		}).Times(0).Return(post, nil)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts/:postId")
		c.SetParamNames("userId", "postId")
		c.SetParamValues(strconv.Itoa(userId), strconv.Itoa(post.Id))

		if assert.NoError(t, handler.updatePost(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("delete - ok", func(t *testing.T) {
		postsService.EXPECT().DeletePost(userId, post.Id).Return(nil)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts/:postId")
		c.SetParamNames("userId", "postId")
		c.SetParamValues(strconv.Itoa(userId), strconv.Itoa(post.Id))

		if assert.NoError(t, handler.deletePost(c)) {
			assert.Equal(t, http.StatusNoContent, rec.Code)
		}
	})

}
