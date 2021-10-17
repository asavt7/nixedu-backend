package server

import (
	"errors"
	mock_service "github.com/asavt7/nixEducation/mocks/pkg/service"
	"github.com/asavt7/nixEducation/pkg/model"
	"github.com/asavt7/nixEducation/pkg/service"
	"github.com/golang/mock/gomock"
	"github.com/kinbiko/jsonassert"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

const (
	username         = "Bret"
	email            = "Sincere@april.biz"
	createduserId    = "15"
	password         = "password"
	createUserRqBody = `{
    "id": 1,
	"password": "` + password + `",
    "name": "Leanne Graham",
    "username": "` + username + `",
    "email": "` + email + `",
    "address": {
      "street": "Kulas Light",
      "suite": "Apt. 556",
      "city": "Gwenborough",
      "zipcode": "92998-3874",
      "geo": {
        "lat": "-37.3159",
        "lng": "81.1496"
      }
    },
    "phone": "1-770-736-8031 x56442",
    "website": "hildegard.org",
    "company": {
      "name": "Romaguera-Crona",
      "catchPhrase": "Multi-layered client-server neural-net",
      "bs": "harness real-time e-markets"
    }
  }`
	createUserRsBodyExpected = `{"id": ` + createduserId + `,"username": "` + username + `", "email": "` + email + `"}`
)

func TestSignUp(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	userService := mock_service.NewMockUserService(controller)

	mockService := &service.Service{
		AuthorizationService: nil,
		UserService:          userService,
		PostService:          nil,
		CommentService:       nil,
	}

	srv := NewApiServer(NewApiHandler(mockService))
	defer srv.Echo.Close()

	t.Run("signUp ok", func(t *testing.T) {
		createduserIdInt, _ := strconv.Atoi(createduserId)
		createdUser := model.User{
			Id:           createduserIdInt,
			Username:     username,
			Email:        email,
			PasswordHash: "hash",
		}
		userService.EXPECT().CreateUser(model.User{
			Id:           1,
			Username:     username,
			Email:        email,
			PasswordHash: "",
		}, password).Return(createdUser, nil)

		req := httptest.NewRequest(echo.POST, "/sign-up", strings.NewReader(createUserRqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		srv.Echo.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)

		assertJsonResponse(t, createUserRsBodyExpected, rec.Body.String())
	})

}

func TestSignIn(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	userService := mock_service.NewMockUserService(controller)
	authService := mock_service.NewMockAuthorizationService(controller)

	mockService := &service.Service{
		AuthorizationService: authService,
		UserService:          userService,
		PostService:          nil,
		CommentService:       nil,
	}

	srv := NewApiServer(NewApiHandler(mockService))
	defer srv.Echo.Close()

	userId := 1
	accessToken := "accessToken"
	refreshToken := "refreshToken"
	user := model.User{
		Id:           userId,
		Username:     username,
		Email:        email,
		PasswordHash: password,
	}

	t.Run("ok", func(t *testing.T) {

		authService.EXPECT().CheckUserCredentials(username, password).Return(user, nil)
		authService.EXPECT().GenerateTokens(userId).Return(accessToken, refreshToken, time.Now(), time.Now(), nil)

		req := httptest.NewRequest(echo.POST, "/sign-in", strings.NewReader(`{"username":"`+username+`","password":"`+password+`"}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		srv.Echo.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.True(t, anyMatch(rec.Header().Values("Set-Cookie"), func(s string) bool {
			return strings.Contains(s, "access-token")
		}))
		assert.True(t, anyMatch(rec.Header().Values("Set-Cookie"), func(s string) bool {
			return strings.Contains(s, "refresh-token")
		}))
		jsonassert.New(t).Assertf(rec.Body.String(), `{			"access-token": "<<PRESENCE>>",				"refresh-token": "<<PRESENCE>>"		}`)
	})

	t.Run("invalid username/password", func(t *testing.T) {

		authService.EXPECT().CheckUserCredentials(username, password).Return(user, errors.New("invalid username/password"))
		//authService.EXPECT().GenerateTokens(userId).Return(accessToken, refreshToken, time.Now(), time.Now(), nil)

		req := httptest.NewRequest(echo.POST, "/sign-in", strings.NewReader(`{"username":"`+username+`","password":"`+password+`"}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		srv.Echo.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.False(t, anyMatch(rec.Header().Values("Set-Cookie"), func(s string) bool {
			return strings.Contains(s, "access-token")
		}))
		assert.False(t, anyMatch(rec.Header().Values("Set-Cookie"), func(s string) bool {
			return strings.Contains(s, "refresh-token")
		}))
	})

}

func anyMatch(s []string, matchFunc func(s string) bool) bool {
	for _, s2 := range s {
		if matchFunc(s2) {
			return true
		}
	}
	return false
}

func assertJsonResponse(t *testing.T, expected, actual string) {
	ja := jsonassert.New(t)
	ja.Assertf(actual, expected)
}
