package server

import (
	// added for swagger support
	_ "github.com/asavt7/nixEducation/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const apiPath = "/api/v1"

// APIServer struct
type APIServer struct {
	Echo    *echo.Echo
	handler *APIHandler
}

// NewAPIServer constructs APIServer
func NewAPIServer(handler *APIHandler) *APIServer {
	s := &APIServer{
		Echo:    echo.New(),
		handler: handler,
	}
	s.initRoutes()
	return s
}

func (srv *APIServer) initRoutes() {
	srv.Echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	srv.Echo.Use(middleware.Recover())

	srv.Echo.GET("/login", srv.handler.loginPage)
	srv.Echo.GET("/oauth/google/login", srv.handler.handleGoogleLogin)
	srv.Echo.GET("/oauth/google/callback", srv.handler.handleGoogleCallback)

	srv.Echo.GET("/swagger/*", echoSwagger.WrapHandler)

	srv.Echo.POST("/sign-in", srv.handler.signIn)
	srv.Echo.POST("/sign-up", srv.handler.signUp)

	srv.Echo.GET("/health", healthCheck)

	API := srv.Echo.Group(apiPath)

	API.Use(parseAccessToken(), srv.handler.tokenRefresherMiddleware)

	usersAPI := API.Group("/users/:userId")

	usersAPI.GET("/posts", srv.handler.getUserPosts)
	usersAPI.POST("/posts", srv.handler.createPost)

	usersAPI.GET("/posts/:postId", srv.handler.getUserPostByID)
	usersAPI.DELETE("/posts/:postId", srv.handler.deletePost)
	usersAPI.PUT("/posts/:postId", srv.handler.updatePost)

	usersAPI.GET("/posts/:postId/comments", srv.handler.getCommentsByPostID)
	usersAPI.POST("/posts/:postId/comments", srv.handler.createComment)

	usersAPI.DELETE("/posts/:postId/comments/:commentId", srv.handler.deleteComment)
	usersAPI.PUT("/posts/:postId/comments/:commentId", srv.handler.updateComment)

}

// Run method run server or fail app
func (srv *APIServer) Run() {

	srv.Echo.Logger.Fatal(srv.Echo.Start(":8080"))

}
