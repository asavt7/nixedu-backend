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

func TestCommentsHandler(t *testing.T) {

	userId := 1
	postId := 1
	commentId := 1
	body := "body"

	comment := model.Comment{
		PostId: postId,
		Id:     commentId,
		UserId: userId,
		Body:   body,
	}
	comments := []model.Comment{comment}

	expectedCommentsJson, err := json.Marshal(comments)
	if err != nil {
		t.Fatal(err)
	}
	expectedCommentJson, err := json.Marshal(comment)
	if err != nil {
		t.Fatal(err)
	}

	controller := gomock.NewController(t)
	defer controller.Finish()
	commentService := mock_service.NewMockCommentService(controller)
	mockService := &service.Service{
		CommentService: commentService,
	}
	handler := NewApiHandler(mockService)

	t.Run("create - ok", func(t *testing.T) {
		commentService.EXPECT().Save(postId, comment).Return(comment, nil)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"body":"`+body+`"}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts/:postId/comments")
		c.SetParamNames("userId", "postId")
		c.SetParamValues(strconv.Itoa(userId), strconv.Itoa(postId))
		c.Set(currentUserId, userId)

		if assert.NoError(t, handler.createComment(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, "/users/1/posts/1/comments/1", rec.Header().Get("Location"))
			jsonassert.New(t).Assertf(rec.Body.String(), string(expectedCommentJson))
		}
	})

	t.Run("create - bad request body", func(t *testing.T) {
		commentService.EXPECT().Save(postId, comment).Times(0).Return(comment, nil)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"body":"`+body+`"}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts/:postId/comments")
		c.SetParamNames("userId", "postId")
		c.SetParamValues(strconv.Itoa(userId), strconv.Itoa(postId))
		c.Set(currentUserId, userId)

		if assert.NoError(t, handler.createComment(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("get all comments by postId", func(t *testing.T) {
		commentService.EXPECT().GetAllByPostId(postId).Return(comments, nil)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts/:postId/comments")
		c.SetParamNames("userId", "postId")
		c.SetParamValues(strconv.Itoa(userId), strconv.Itoa(postId))
		c.Set(currentUserId, 1111)

		if assert.NoError(t, handler.getCommentsByPostId(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			jsonassert.New(t).Assertf(rec.Body.String(), string(expectedCommentsJson))
		}
	})

	t.Run("get all comments by postId - user not found", func(t *testing.T) {
		commentService.EXPECT().GetAllByPostId(postId).Return(comments, service.UserNotFoundErr{Id: userId})

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts/:postId/comments")
		c.SetParamNames("userId", "postId")
		c.SetParamValues(strconv.Itoa(userId), strconv.Itoa(postId))
		c.Set(currentUserId, 1111)

		if assert.NoError(t, handler.getCommentsByPostId(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}
	})

	t.Run("update - ok", func(t *testing.T) {
		commentService.EXPECT().Update(userId, commentId, model.UpdateComment{Body: &body}).Return(comment, nil)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(`{"body":"`+body+`"}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts/:postId/comments/:commentId")
		c.SetParamNames("userId", "postId", "commentId")
		c.SetParamValues(strconv.Itoa(userId), strconv.Itoa(postId), strconv.Itoa(commentId))
		c.Set(currentUserId, userId)

		if assert.NoError(t, handler.updateComment(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			jsonassert.New(t).Assertf(rec.Body.String(), string(expectedCommentJson))
		}
	})

	t.Run("update - bad request - empty body", func(t *testing.T) {
		commentService.EXPECT().Update(userId, commentId, model.UpdateComment{Body: &body}).Times(0).Return(comment, nil)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(`{}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts/:postId/comments/:commentId")
		c.SetParamNames("userId", "postId", "commentId")
		c.SetParamValues(strconv.Itoa(userId), strconv.Itoa(postId), strconv.Itoa(commentId))
		c.Set(currentUserId, userId)

		if assert.NoError(t, handler.updateComment(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("update - unauthorized", func(t *testing.T) {
		commentService.EXPECT().Update(userId, commentId, model.UpdateComment{Body: &body}).Times(0).Return(comment, nil)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(`{"body":"`+body+`"}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts/:postId/comments/:commentId")
		c.SetParamNames("userId", "postId", "commentId")
		c.SetParamValues(strconv.Itoa(userId), strconv.Itoa(postId), strconv.Itoa(commentId))
		c.Set(currentUserId, 333333)

		if assert.NoError(t, handler.updateComment(c)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})

	t.Run("delete - ok", func(t *testing.T) {
		commentService.EXPECT().Delete(userId, commentId).Return(nil)

		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", strings.NewReader(`{"body":"`+body+`"}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts/:postId/comments/:commentId")
		c.SetParamNames("userId", "postId", "commentId")
		c.SetParamValues(strconv.Itoa(userId), strconv.Itoa(postId), strconv.Itoa(commentId))
		c.Set(currentUserId, userId)

		if assert.NoError(t, handler.deleteComment(c)) {
			assert.Equal(t, http.StatusNoContent, rec.Code)
		}
	})

	t.Run("delete - unauthorized", func(t *testing.T) {
		commentService.EXPECT().Delete(userId, commentId).Times(0).Return(nil)

		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", strings.NewReader(`{"body":"`+body+`"}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts/:postId/comments/:commentId")
		c.SetParamNames("userId", "postId", "commentId")
		c.SetParamValues(strconv.Itoa(userId), strconv.Itoa(postId), strconv.Itoa(commentId))
		c.Set(currentUserId, 3333333)

		if assert.NoError(t, handler.deleteComment(c)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})

}
