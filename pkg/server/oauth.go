package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/asavt7/nixEducation/pkg/configs"
	"github.com/asavt7/nixEducation/pkg/model"
	"github.com/asavt7/nixEducation/pkg/service"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
)

var (
	oauthStateString  = uuid.NewString()
	googleOauthConfig *oauth2.Config
)

func init() {
	googleOauthConfig = configs.InitGoogleOAuthConfigs()
}

func (h *ApiHandler) handleGoogleLogin(context echo.Context) error {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	return context.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *ApiHandler) handleGoogleCallback(c echo.Context) error {
	content, err := getUserInfo(c.FormValue("state"), c.FormValue("code"))
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	googleUser := &struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}{}
	err = json.Unmarshal(content, googleUser)
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	gotUser, err := h.getUserByEmailOrCreateIfNotExists(googleUser.Email, googleUser.Name)
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	_, _, err = h.generateTokensAndSetCookies(gotUser.Id, c)
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/swagger/")
}

func (h *ApiHandler) getUserByEmailOrCreateIfNotExists(email, username string) (model.User, error) {
	u, err := h.service.UserService.GetUserByEmail(email)
	if err != nil {
		switch err.(type) {
		case service.UserNotFoundErr:
			return h.service.UserService.CreateUser(model.User{
				Username: email,
				Email:    username,
			}, "")
		default:
			return model.User{}, err
		}
	}
	return u, nil
}

func getUserInfo(state string, code string) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return contents, nil
}
