package server

import (
	"github.com/asavt7/nixEducation/pkg/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// createComment godoc
// @Tags comments
// @Summary createComment
// @Description createComment
// @ID createComment
// @Accept  json,xml
// @Produce  json,xml
// @Param userId path int true "userId"
// @Param postId path int true "postId"
// @Param comment body model.Comment true "post"
// @Success 201 {object} model.Comment
// @Failure 400 {object} message
// @Failure 500 {object} message
// @Router /api/v1/users/{userId}/posts/{postId}/comments [post]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (h *ApiHandler) createComment(context echo.Context) error {

	postId := context.Param("postId")
	postIDInt, err := strconv.Atoi(postId)
	if err != nil {
		return response(http.StatusBadRequest, "missing or incorrect userId param, expected int", context)
	}

	newComment := new(model.Comment)
	if err := context.Bind(newComment); err != nil {
		return response(http.StatusBadRequest, err.Error(), context)
	}

	newComment.UserId = context.Get(currentUserId).(int)
	newComment.PostId = postIDInt

	if err := h.validator.Struct(newComment); err != nil {
		return response(http.StatusBadRequest, err.Error(), context)
	}

	post, err := h.service.CommentService.Save(postIDInt, *newComment)
	if err != nil {
		switch err.(type) {
		case model.UserNotFoundErr, model.PostNotFoundErr:
			return response(http.StatusNotFound, message{Message: err.Error()}, context)
		default:
			return response(http.StatusInternalServerError, message{Message: err.Error()}, context)
		}
	}
	return response(http.StatusCreated, post, context)
}

// getCommentsByPostId godoc
// @Tags comments
// @Summary getCommentsByPostId
// @Description getCommentsByPostId
// @ID getCommentsByPostId
// @Accept  json,xml
// @Produce  json,xml
// @Param userId path int true "userId"
// @Param postId path int true "postId"
// @Success 200 {object} []model.Comment
// @Failure 400 {object} message
// @Failure 500 {object} message
// @Router /api/v1/users/{userId}/posts/{postId}/comments [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (h *ApiHandler) getCommentsByPostId(context echo.Context) error {
	postId := context.Param("postId")
	postIDInt, err := strconv.Atoi(postId)
	if err != nil {
		return response(http.StatusBadRequest, "missing or incorrect postId param, expected int", context)
	}

	post, err := h.service.CommentService.GetAllByPostId(postIDInt)
	if err != nil {
		switch err.(type) {
		case model.UserNotFoundErr, model.PostNotFoundErr:
			return response(http.StatusNotFound, message{Message: err.Error()}, context)
		default:
			return response(http.StatusInternalServerError, message{Message: err.Error()}, context)
		}
	}
	return response(http.StatusOK, post, context)
}

// deleteComment godoc
// @Tags comments
// @Summary deleteComment
// @Description deleteComment
// @ID deleteComment
// @Accept  json,xml
// @Produce  json,xml
// @Param userId path int true "userId"
// @Param postId path int true "postId"
// @Param commentId path int true "commentId"
// @Success 204
// @Failure 400 {object} message
// @Failure 500 {object} message
// @Router /api/v1/users/{userId}/posts/{postId}/comments/{commentId} [delete]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (h *ApiHandler) deleteComment(context echo.Context) error {

	commentId := context.Param("commentId")
	commentIDInt, err := strconv.Atoi(commentId)
	if err != nil {
		return response(http.StatusBadRequest, "missing or incorrect commentId param, expected int", context)
	}

	currentUser := context.Get(currentUserId).(int)

	if err := h.service.CommentService.Delete(currentUser, commentIDInt); err != nil {
		switch err.(type) {
		case model.CommentNotFoundErr:
			return response(http.StatusNotFound, message{Message: err.Error()}, context)
		case model.UserHasNoAccessToChangeComment:
			return response(http.StatusUnauthorized, message{Message: err.Error()}, context)
		default:
			return response(http.StatusInternalServerError, message{Message: err.Error()}, context)
		}
	}
	return context.NoContent(http.StatusNoContent)

}

// updateComment godoc
// @Tags comments
// @Summary updateComment
// @Description updateComment
// @ID updateComment
// @Accept  json,xml
// @Produce  json,xml
// @Param userId path int true "userId"
// @Param postId path int true "postId"
// @Param commentId path int true "commentId"
// @Param comment body model.UpdateComment true "post"
// @Success 200 {object} model.Comment
// @Failure 400 {object} message
// @Failure 500 {object} message
// @Router /api/v1/users/{userId}/posts/{postId}/comments/{commentId} [put]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (h *ApiHandler) updateComment(context echo.Context) error {

	commentId := context.Param("commentId")
	commentIDInt, err := strconv.Atoi(commentId)
	if err != nil {
		return response(http.StatusBadRequest, "missing or incorrect commentId param, expected int", context)
	}

	currentUser := context.Get(currentUserId).(int)

	updateInput := new(model.UpdateComment)
	if err := context.Bind(updateInput); err != nil {
		return response(http.StatusBadRequest, err.Error(), context)
	}
	if err := h.validator.Struct(updateInput); err != nil {
		return response(http.StatusBadRequest, err.Error(), context)
	}

	post, err := h.service.CommentService.Update(currentUser, commentIDInt, *updateInput)
	if err != nil {
		switch err.(type) {
		case model.CommentNotFoundErr:
			return response(http.StatusNotFound, message{Message: err.Error()}, context)
		case model.UserHasNoAccessToChangeComment:
			return response(http.StatusUnauthorized, message{Message: err.Error()}, context)
		default:
			return response(http.StatusInternalServerError, message{Message: err.Error()}, context)
		}
	}
	return response(http.StatusOK, post, context)

}
