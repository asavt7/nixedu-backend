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

	userID := 1
	postID := 1
	commentID := 1
	body := "body"

	comment := model.Comment{
		PostID: postID,
		ID:     commentID,
		UserID: userID,
		Body:   body,
	}
	comments := []model.Comment{comment}

	expectedCommentsJSON, err := json.Marshal(comments)
	if err != nil {
		t.Fatal(err)
	}
	expectedCommentJSON, err := json.Marshal(comment)
	if err != nil {
		t.Fatal(err)
	}

	controller := gomock.NewController(t)
	defer controller.Finish()
	commentService := mock_service.NewMockCommentService(controller)
	mockService := &service.Service{
		CommentService: commentService,
	}
	handler := NewAPIHandler(mockService)

	t.Run("create - ok", func(t *testing.T) {
		commentService.EXPECT().Save(model.Comment{
			PostID: postID,
			ID:     0,
			UserID: userID,
			Body:   body,
		}).Return(comment, nil)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"body":"`+body+`"}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts/:postId/comments")
		c.SetParamNames("userId", "postId")
		c.SetParamValues(strconv.Itoa(userID), strconv.Itoa(postID))
		c.Set(currentUserID, userID)

		if assert.NoError(t, handler.createComment(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			jsonassert.New(t).Assertf(rec.Body.String(), string(expectedCommentJSON))
		}
	})

	t.Run("create - bad request body", func(t *testing.T) {
		commentService.EXPECT().Save(comment).Times(0).Return(comment, nil)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts/:postId/comments")
		c.SetParamNames("userId", "postId")
		c.SetParamValues(strconv.Itoa(userID), strconv.Itoa(postID))
		c.Set(currentUserID, userID)

		if assert.NoError(t, handler.createComment(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("get all comments by postID", func(t *testing.T) {
		commentService.EXPECT().GetAllByPostID(postID).Return(comments, nil)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts/:postId/comments")
		c.SetParamNames("userId", "postId")
		c.SetParamValues(strconv.Itoa(userID), strconv.Itoa(postID))
		c.Set(currentUserID, 1111)

		if assert.NoError(t, handler.getCommentsByPostID(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			jsonassert.New(t).Assertf(rec.Body.String(), string(expectedCommentsJSON))
		}
	})

	t.Run("get all comments by postID - user not found", func(t *testing.T) {
		commentService.EXPECT().GetAllByPostID(postID).Return(comments, model.UserNotFoundErr{ID: userID})

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts/:postId/comments")
		c.SetParamNames("userId", "postId")
		c.SetParamValues(strconv.Itoa(userID), strconv.Itoa(postID))
		c.Set(currentUserID, 1111)

		if assert.NoError(t, handler.getCommentsByPostID(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}
	})

	t.Run("update - ok", func(t *testing.T) {
		commentService.EXPECT().Update(userID, commentID, model.UpdateComment{Body: &body}).Return(comment, nil)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(`{"body":"`+body+`"}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts/:postId/comments/:commentId")
		c.SetParamNames("userId", "postId", "commentId")
		c.SetParamValues(strconv.Itoa(userID), strconv.Itoa(postID), strconv.Itoa(commentID))
		c.Set(currentUserID, userID)

		if assert.NoError(t, handler.updateComment(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			jsonassert.New(t).Assertf(rec.Body.String(), string(expectedCommentJSON))
		}
	})

	t.Run("update - bad request - empty body", func(t *testing.T) {
		commentService.EXPECT().Update(userID, commentID, model.UpdateComment{Body: &body}).Times(0).Return(comment, nil)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(`{}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts/:postId/comments/:commentId")
		c.SetParamNames("userId", "postId", "commentId")
		c.SetParamValues(strconv.Itoa(userID), strconv.Itoa(postID), strconv.Itoa(commentID))
		c.Set(currentUserID, userID)

		if assert.NoError(t, handler.updateComment(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("update - unauthorized", func(t *testing.T) {
		commentService.EXPECT().Update(userID, commentID, model.UpdateComment{Body: &body}).Return(comment, model.UserHasNoAccessToChangeComment{
			UserID:    userID,
			CommentID: commentID,
		})

		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(`{"body":"`+body+`"}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts/:postId/comments/:commentId")
		c.SetParamNames("userId", "postId", "commentId")
		c.SetParamValues(strconv.Itoa(userID), strconv.Itoa(postID), strconv.Itoa(commentID))
		c.Set(currentUserID, userID)

		if assert.NoError(t, handler.updateComment(c)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})

	t.Run("delete - ok", func(t *testing.T) {
		commentService.EXPECT().Delete(userID, commentID).Return(nil)

		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", strings.NewReader(`{"body":"`+body+`"}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts/:postId/comments/:commentId")
		c.SetParamNames("userId", "postId", "commentId")
		c.SetParamValues(strconv.Itoa(userID), strconv.Itoa(postID), strconv.Itoa(commentID))
		c.Set(currentUserID, userID)

		if assert.NoError(t, handler.deleteComment(c)) {
			assert.Equal(t, http.StatusNoContent, rec.Code)
		}
	})

	t.Run("delete - unauthorized", func(t *testing.T) {
		commentService.EXPECT().Delete(userID, commentID).Return(model.UserHasNoAccessToChangeComment{
			UserID:    userID,
			CommentID: commentID,
		})

		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", strings.NewReader(`{"body":"`+body+`"}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:userId/posts/:postId/comments/:commentId")
		c.SetParamNames("userId", "postId", "commentId")
		c.SetParamValues(strconv.Itoa(userID), strconv.Itoa(postID), strconv.Itoa(commentID))
		c.Set(currentUserID, userID)

		if assert.NoError(t, handler.deleteComment(c)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})

}
