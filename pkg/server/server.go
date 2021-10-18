package server

import (
	_ "github.com/asavt7/nixEducation/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const ApiPath = "/api/v1"

type ApiServer struct {
	Echo    *echo.Echo
	handler *ApiHandler
}

func NewApiServer(handler *ApiHandler) *ApiServer {
	s := &ApiServer{
		Echo:    echo.New(),
		handler: handler,
	}
	s.InitRoutes()
	return s
}

func (srv *ApiServer) InitRoutes() {
	srv.Echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	srv.Echo.Use(middleware.Recover())

	srv.Echo.GET("/swagger/*", echoSwagger.WrapHandler)

	srv.Echo.POST("/sign-in", srv.handler.signIn)
	srv.Echo.POST("/sign-up", srv.handler.signUp)

	srv.Echo.GET("/health", healthCheck)

	api := srv.Echo.Group(ApiPath)

	api.Use(parseAccessToken(), srv.handler.TokenRefresherMiddleware)

	usersApi := api.Group("/users/:userId")

	usersApi.GET("/posts", srv.handler.getUserPosts)
	usersApi.POST("/posts", srv.handler.createPost)

	usersApi.GET("/posts/:postId", srv.handler.getUserPostById)
	usersApi.DELETE("/posts/:postId", srv.handler.deletePost)
	usersApi.PUT("/posts/:postId", srv.handler.updatePost)

	usersApi.GET("/posts/:postId/comments", srv.handler.getCommentsByPostId)
	usersApi.POST("/posts/:postId/comments", srv.handler.createComment)

	usersApi.DELETE("/posts/:postId/comments/:commentId", srv.handler.deleteComment)
	usersApi.PUT("/posts/:postId/comments/:commentId", srv.handler.updateComment)

}

func (srv *ApiServer) Run() {

	srv.Echo.Logger.Fatal(srv.Echo.Start(":8080"))

}
