package server

import (
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

func assertJsonResponse(t *testing.T, expected, actual string) {
	ja := jsonassert.New(t)
	ja.Assertf(actual, expected)
}
