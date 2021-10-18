package server

import "github.com/labstack/echo/v4"

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
// @Failure 400 {object} Message
// @Failure 500 {object} Message
// @Router /api/v1/users/{userId}/posts/{postId}/comments [post]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (h *ApiHandler) createComment(context echo.Context) error {
	return nil
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
// @Failure 400 {object} Message
// @Failure 500 {object} Message
// @Router /api/v1/users/{userId}/posts/{postId}/comments [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (h *ApiHandler) getCommentsByPostId(context echo.Context) error {
	return nil

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
// @Failure 400 {object} Message
// @Failure 500 {object} Message
// @Router /api/v1/users/{userId}/posts/{postId}/comments/{commentId} [delete]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (h *ApiHandler) deleteComment(context echo.Context) error {
	return nil

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
// @Failure 400 {object} Message
// @Failure 500 {object} Message
// @Router /api/v1/users/{userId}/posts/{postId}/comments/{commentId} [put]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (h *ApiHandler) updateComment(context echo.Context) error {
	return nil

}
