package server

import (
	"github.com/asavt7/nixEducation/pkg/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// getUserPosts godoc
// @Tags posts
// @Summary getUserPosts
// @Description get posts by userId
// @ID getUserPosts
// @Accept  json,xml
// @Produce  json,xml
// @Param userId path int true "userId"
// @Success 200 {object} []model.Post
// @Failure 404 {object} message
// @Failure 500 {object} message
// @Router /api/v1/users/{userId}/posts [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (h *APIHandler) getUserPosts(context echo.Context) error {
	userID := context.Param("userId")
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return response(http.StatusBadRequest, "missing or incorrect userId param, expected int", context)
	}

	posts, err := h.service.PostService.GetAllByUserID(userIDInt)
	if err != nil {
		switch err.(type) {
		case model.UserNotFoundErr:
			return response(http.StatusNotFound, message{Message: err.Error()}, context)
		default:
			return response(http.StatusInternalServerError, message{Message: err.Error()}, context)
		}
	}
	return response(http.StatusOK, posts, context)
}

// createPost godoc
// @Tags posts
// @Summary createPost
// @Description createPost
// @ID createPost
// @Accept  json,xml
// @Produce  json,xml
// @Param userId path int true "userId"
// @Param post body model.Post true "post"
// @Success 201 {object} model.Post
// @Header 201 {string} Location "/api/v1/users/{userId}/posts/{postId}"
// @Failure 400 {object} message
// @Failure 500 {object} message
// @Router /api/v1/users/{userId}/posts [post]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (h *APIHandler) createPost(context echo.Context) error {

	userID := context.Param("userId")
	userIDInt, err := strconv.Atoi(userID)
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

	currentUser := context.Get(currentUserID).(int)
	if currentUser != userIDInt {
		return response(http.StatusUnauthorized, "unauthorized", context)
	}

	post, err := h.service.PostService.Save(userIDInt, *newPost)
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

// getUserPostByID godoc
// @Tags posts
// @Summary getUserPostByID
// @Description getUserPostByID
// @ID getUserPostByID
// @Accept  json,xml
// @Produce  json,xml
// @Param userId path int true "userId"
// @Param userId path int true "postId"
// @Success 200 {object} model.Post
// @Failure 404 {object} message
// @Failure 500 {object} message
// @Router /api/v1/users/{userId}/posts/{postId} [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (h *APIHandler) getUserPostByID(context echo.Context) error {

	userID := context.Param("userId")
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return response(http.StatusBadRequest, "missing or incorrect userId param, expected int", context)
	}
	postID := context.Param("postId")
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		return response(http.StatusBadRequest, "missing or incorrect postId param, expected int", context)
	}

	post, err := h.service.PostService.GetByUserIDAndPostID(userIDInt, postIDInt)
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

// deletePost godoc
// @Tags posts
// @Summary deletePost
// @Description deletePost
// @ID deletePost
// @Accept  json,xml
// @Produce  json,xml
// @Param userId path int true "userId"
// @Param userId path int true "postId"
// @Success 204 {object} model.Post
// @Failure 404 {object} message
// @Failure 500 {object} message
// @Router /api/v1/users/{userId}/posts/{postId} [delete]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (h *APIHandler) deletePost(context echo.Context) error {
	userID := context.Param("userId")
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return response(http.StatusBadRequest, "missing or incorrect userId param, expected int", context)
	}
	postID := context.Param("postId")
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		return response(http.StatusBadRequest, "missing or incorrect postId param, expected int", context)
	}

	currentUser := context.Get(currentUserID).(int)
	if currentUser != userIDInt {
		return response(http.StatusUnauthorized, "unauthorized", context)
	}

	if err := h.service.PostService.DeletePost(userIDInt, postIDInt); err != nil {
		switch err.(type) {
		case model.UserNotFoundErr, model.PostNotFoundErr:
			return response(http.StatusNotFound, message{Message: err.Error()}, context)
		default:
			return response(http.StatusInternalServerError, message{Message: err.Error()}, context)
		}
	}
	return context.NoContent(http.StatusNoContent)

}

// updatePost godoc
// @Tags posts
// @Summary updatePost
// @Description updatePost
// @ID updatePost
// @Accept  json,xml
// @Produce  json,xml
// @Param userId path int true "userId"
// @Param userId path int true "postId"
// @Param post body model.UpdatePost true "post"
// @Success 200 {object} model.Post
// @Failure 404 {object} message
// @Failure 500 {object} message
// @Router /api/v1/users/{userId}/posts/{postId} [put]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (h *APIHandler) updatePost(context echo.Context) error {
	userID := context.Param("userId")
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return response(http.StatusBadRequest, "missing or incorrect userID param, expected int", context)
	}
	postID := context.Param("postId")
	postIDInt, err := strconv.Atoi(postID)
	if err != nil {
		return response(http.StatusBadRequest, "missing or incorrect postID param, expected int", context)
	}

	currentUser := context.Get(currentUserID).(int)
	if currentUser != userIDInt {
		return response(http.StatusUnauthorized, "unauthorized", context)
	}

	updatePostInput := new(model.UpdatePost)
	if err := context.Bind(updatePostInput); err != nil {
		return response(http.StatusBadRequest, err.Error(), context)
	}
	if err := h.validator.Struct(updatePostInput); err != nil {
		return response(http.StatusBadRequest, err.Error(), context)
	}

	post, err := h.service.PostService.Update(userIDInt, postIDInt, *updatePostInput)
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
