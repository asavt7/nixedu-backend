package server

import (
	"github.com/asavt7/nixedu/backend/pkg/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

const refreshTokenCookieName = "refresh-token"
const accessTokenCookieName = "access-token"
const currentUserID = "currentUserID"

type signInUserInput struct {
	Password string `json:"password" xml:"password" validate:"required"`
	Username string `json:"username" xml:"username" validate:"required"`
}

type signInResponse struct {
	AccessToken  string `json:"access-token" xml:"access-token"`
	RefreshToken string `json:"refresh-token" xml:"refresh-token"`
}

// signIn godoc
// @Tags auth
// @Summary signIn
// @Description signIn and get access and refresh tokens
// @ID signIn
// @Accept  json,xml
// @Produce  json,xml
// @Param signInUserInput body signInUserInput true "body"
// @Success 200 {object} signInResponse
// @Failure 400 {object} message
// @Failure 500 {object} message
// @Router /sign-in [post]
func (h *APIHandler) signIn(c echo.Context) error {
	u := new(signInUserInput)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if len(u.Password) == 0 || len(u.Username) == 0 {
		u.Username = c.FormValue("username")
		u.Password = c.FormValue("password")
	}

	err := h.validator.Struct(u)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := h.service.AuthorizationService.CheckUserCredentials(u.Username, u.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Password or Username is incorrect")
	}

	accessToken, refreshToken, err := h.generateTokensAndSetCookies(user.ID, c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Token is incorrect")
	}

	return response(http.StatusOK, signInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, c)
}

func (h *APIHandler) generateTokensAndSetCookies(userID int, c echo.Context) (accessToken, refreshToken string, err error) {
	accessToken, refreshToken, accessExp, refreshExp, err := h.service.AuthorizationService.GenerateTokens(userID)
	if err != nil {
		return accessToken, refreshToken, err
	}

	h.setTokenCookie(accessTokenCookieName, accessToken, accessExp, c)
	h.setTokenCookie(refreshTokenCookieName, refreshToken, refreshExp, c)

	return accessToken, refreshToken, nil
}

func (h *APIHandler) setTokenCookie(name, token string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expiration
	cookie.Path = apiPath
	// Http-only helps mitigate the risk of client side script accessing the protected cookie.
	cookie.HttpOnly = true

	c.SetCookie(cookie)
}

func (h *APIHandler) setUserCookie(userID string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "userId"
	cookie.Value = userID
	cookie.Expires = expiration
	cookie.Path = "/"
	c.SetCookie(cookie)
}

type signUpUserInput struct {
	Password string `json:"password" xml:"password" validate:"required"`
	model.User
}

// signUp godoc
// @Tags auth
// @Summary signUp
// @Description signUp new user
// @ID signUp
// @Accept  json,xml
// @Produce  json,xml
// @Param signUpUserInput body signUpUserInput true "a body"
// @Success 200 {object} model.User
// @Failure 400 {object} message
// @Failure 500 {object} message
// @Router /sign-up [post]
func (h *APIHandler) signUp(c echo.Context) error {
	u := new(signUpUserInput)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.validator.Struct(u)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	createdUser, err := h.service.UserService.CreateUser(u.User, u.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return response(http.StatusCreated, createdUser, c)
}

func (h *APIHandler) loginPage(context echo.Context) error {
	var htmlLoginForm = `<html>
<body>
<center> <h1> Login Form </h1> </center>   
    <form action="/sign-in" method="post">
		  <label for="username">Username:</label>
		  <input type="text" id="username" name="username"><br><br>
		  <label for="password">Password:</label>
		  <input type="password" id="password" name="password"><br><br>
		  <input type="submit" value="Submit">
	</form>

	<a href="/oauth/google/login">Google Log In</a>
</body>
</html>`
	return context.HTML(200, htmlLoginForm)
}
