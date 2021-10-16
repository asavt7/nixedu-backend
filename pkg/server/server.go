package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

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
	srv.Echo.Use(middleware.Logger())
	srv.Echo.Use(middleware.Recover())

	srv.Echo.POST("/sign-in", srv.handler.signIn)
	srv.Echo.POST("/sign-up", srv.handler.signUp)

	api := srv.Echo.Group("/api/v1")

	api.Use(parseAccessToken(), srv.handler.TokenRefresherMiddleware)

	api.GET("/health", func(context echo.Context) error {
		return context.NoContent(200)
	})

	usersApi := api.Group("/users/:userId")

	usersApi.GET("/posts", srv.handler.getUserPosts)
	usersApi.POST("/posts", srv.handler.createPost)

	usersApi.GET("/posts/:postId", srv.handler.getUserPostById)
	usersApi.DELETE("/posts/:postId", srv.handler.deletePost)
	usersApi.PUT("/posts/:postId", srv.handler.changePost)

	usersApi.GET("/posts/:postId/comments", srv.handler.getComments)
	usersApi.POST("/posts/:postId/comments", srv.handler.createComment)

	usersApi.GET("/posts/:postId/comments/:commentId", srv.handler.getCommentById)
	usersApi.DELETE("/posts/:postId/comments/:commentId", srv.handler.deleteComment)
	usersApi.PUT("/posts/:postId/comments/:commentId", srv.handler.changeComment)

}

func (srv *ApiServer) Run() {

	srv.Echo.Logger.Fatal(srv.Echo.Start(":8080"))

}
