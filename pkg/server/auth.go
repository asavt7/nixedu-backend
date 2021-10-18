package server

import (
	"github.com/asavt7/nixEducation/pkg/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"time"
)

const refreshTokenCookieName = "refresh-token"
const accessTokenCookieName = "access-token"
const currentUserId = "currentUserId"

type signInUserInput struct {
	Password string `json:"password" xml:"password"`
	Username string `json:"username" xml:"username"`
}

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

	return response(http.StatusOK, map[string]string{
		"access-token":  accessToken,
		"refresh-token": refreshToken,
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
	cookie.Path = "/"
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

func response(status int, body interface{}, c echo.Context) error {
	ctype := c.Request().Header.Get(echo.HeaderContentType)
	acceptType := c.Request().Header.Get(echo.HeaderAccept)
	if len(acceptType) == 0 {
		acceptType = ctype
	}
	switch {
	case strings.Contains(acceptType, echo.MIMEApplicationJSON):
		return c.JSON(status, body)
	case strings.Contains(acceptType, echo.MIMEApplicationXML), strings.Contains(acceptType, echo.MIMETextXML):
		return c.XML(status, body)
	default:
		return c.JSON(status, body)
	}
}
