package server

import (
	"github.com/asavt7/nixEducation/pkg/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

const refreshTokenCookieName = "refresh-token"
const accessTokenCookieName = "access-token"
const currentUserId = "currentUserId"

type signInUserInput struct {
	Password string `json:"password" xml:"password"`
	Username string `json:"username" xml:"username"`
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
// @Failure 400 {object} Message
// @Failure 500 {object} Message
// @Router /sign-in [post]
func (h *ApiHandler) signIn(c echo.Context) error {
	u := new(signInUserInput)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := h.service.AuthorizationService.CheckUserCredentials(u.Username, u.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Password or Username is incorrect")
	}

	accessToken, refreshToken, err := h.generateTokensAndSetCookies(user.Id, c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Token is incorrect")
	}

	return response(http.StatusOK, signInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, c)
}

func (h *ApiHandler) generateTokensAndSetCookies(userId int, c echo.Context) (accessToken, refreshToken string, err error) {
	accessToken, refreshToken, accessExp, refreshExp, err := h.service.AuthorizationService.GenerateTokens(userId)
	if err != nil {
		return accessToken, refreshToken, err
	}

	h.setTokenCookie(accessTokenCookieName, accessToken, accessExp, c)
	h.setTokenCookie(refreshTokenCookieName, refreshToken, refreshExp, c)

	return accessToken, refreshToken, nil
}

func (h *ApiHandler) setTokenCookie(name, token string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expiration
	cookie.Path = ApiPath
	// Http-only helps mitigate the risk of client side script accessing the protected cookie.
	cookie.HttpOnly = true

	c.SetCookie(cookie)
}

func (h *ApiHandler) setUserCookie(userId string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "userId"
	cookie.Value = userId
	cookie.Expires = expiration
	cookie.Path = "/"
	c.SetCookie(cookie)
}

type signUpUserInput struct {
	Password string `json:"password" xml:"password"`
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
// @Failure 400 {object} Message
// @Failure 500 {object} Message
// @Router /sign-up [post]
func (h *ApiHandler) signUp(c echo.Context) error {
	u := new(signUpUserInput)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	createdUser, err := h.service.UserService.CreateUser(u.User, u.Password)
	if err != nil {
		return err
	}

	return response(http.StatusCreated, createdUser, c)
}
