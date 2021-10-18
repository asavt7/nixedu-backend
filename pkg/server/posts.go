package server

import (
	"github.com/asavt7/nixEducation/pkg/model"
	"github.com/asavt7/nixEducation/pkg/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *ApiHandler) getUserPosts(context echo.Context) error {
	userId := context.Param("userId")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return response(http.StatusBadRequest, "missing or incorrect userId param, expected int", context)
	}

	posts, err := h.service.PostService.GetAllByUserId(userIdInt)
	if err != nil {
		switch err.(type) {
		case service.UserNotFoundErr:
			return response(http.StatusNotFound, Message{Message: err.Error()}, context)
		default:
			return response(http.StatusInternalServerError, Message{Message: err.Error()}, context)
		}
	}
	return response(http.StatusOK, posts, context)
}

func (h *ApiHandler) createPost(context echo.Context) error {

	userId := context.Param("userId")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return response(http.StatusBadRequest, "missing or incorrect userId param, expected int", context)
	}

	newPost := new(model.Post)
	if err := context.Bind(newPost); err != nil {
		return response(http.StatusBadRequest, err.Error(), context)
	}
	if err := h.validator.Struct(newPost); err != nil {
		return response(http.StatusBadRequest, err.Error(), context)
	}

	currentUser := context.Get(currentUserId).(int)
	if currentUser != userIdInt {
		return response(http.StatusUnauthorized, "unauthorized", context)
	}

	post, err := h.service.PostService.Save(userIdInt, *newPost)
	if err != nil {
		switch err.(type) {
		case service.UserNotFoundErr, service.PostNotFoundErr:
			return response(http.StatusNotFound, Message{Message: err.Error()}, context)
		default:
			return response(http.StatusInternalServerError, Message{Message: err.Error()}, context)
		}
	}
	return response(http.StatusCreated, post, context)

}

func (h *ApiHandler) getUserPostById(context echo.Context) error {

	userId := context.Param("userId")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return response(http.StatusBadRequest, "missing or incorrect userId param, expected int", context)
	}
	postId := context.Param("postId")
	postIdInt, err := strconv.Atoi(postId)
	if err != nil {
		return response(http.StatusBadRequest, "missing or incorrect postId param, expected int", context)
	}

	post, err := h.service.PostService.GetByUserIdAndPostId(userIdInt, postIdInt)
	if err != nil {
		switch err.(type) {
		case service.UserNotFoundErr, service.PostNotFoundErr:
			return response(http.StatusNotFound, Message{Message: err.Error()}, context)
		default:
			return response(http.StatusInternalServerError, Message{Message: err.Error()}, context)
		}
	}
	return response(http.StatusOK, post, context)

}

func (h *ApiHandler) deletePost(context echo.Context) error {
	userId := context.Param("userId")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return response(http.StatusBadRequest, "missing or incorrect userId param, expected int", context)
	}
	postId := context.Param("postId")
	postIdInt, err := strconv.Atoi(postId)
	if err != nil {
		return response(http.StatusBadRequest, "missing or incorrect postId param, expected int", context)
	}

	currentUser := context.Get(currentUserId).(int)
	if currentUser != userIdInt {
		return response(http.StatusUnauthorized, "unauthorized", context)
	}

	if err := h.service.PostService.DeletePost(userIdInt, postIdInt); err != nil {
		switch err.(type) {
		case service.UserNotFoundErr, service.PostNotFoundErr:
			return response(http.StatusNotFound, Message{Message: err.Error()}, context)
		default:
			return response(http.StatusInternalServerError, Message{Message: err.Error()}, context)
		}
	}
	return context.NoContent(http.StatusNoContent)

}

func (h *ApiHandler) updatePost(context echo.Context) error {
	userId := context.Param("userId")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return response(http.StatusBadRequest, "missing or incorrect userId param, expected int", context)
	}
	postId := context.Param("postId")
	postIdInt, err := strconv.Atoi(postId)
	if err != nil {
		return response(http.StatusBadRequest, "missing or incorrect postId param, expected int", context)
	}

	currentUser := context.Get(currentUserId).(int)
	if currentUser != userIdInt {
		return response(http.StatusUnauthorized, "unauthorized", context)
	}

	updatePostInput := new(model.UpdatePost)
	if err := context.Bind(updatePostInput); err != nil {
		return response(http.StatusBadRequest, err.Error(), context)
	}
	if err := h.validator.Struct(updatePostInput); err != nil {
		return response(http.StatusBadRequest, err.Error(), context)
	}

	post, err := h.service.PostService.Update(userIdInt, postIdInt, *updatePostInput)
	if err != nil {
		switch err.(type) {
		case service.UserNotFoundErr, service.PostNotFoundErr:
			return response(http.StatusNotFound, Message{Message: err.Error()}, context)
		default:
			return response(http.StatusInternalServerError, Message{Message: err.Error()}, context)
		}
	}
	return response(http.StatusOK, post, context)
}
