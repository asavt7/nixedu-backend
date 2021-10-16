package server

import (
	"github.com/asavt7/nixEducation/pkg/model"
	"github.com/asavt7/nixEducation/pkg/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSignUp(t *testing.T) {

	mockUSerService := &service.MockUserService{
		CreateUserFunc: nil,
		GetUserFunc:    nil,
	}

	mockService := &service.Service{
		AuthorizationService: nil,
		UserService:          mockUSerService,
		PostService:          nil,
		CommentService:       nil,
	}

	srv := NewApiServer(NewApiHandler(mockService))

	t.Run("signUp ok", func(t *testing.T) {
		calls := 0
		var userToCreate model.User
		var passwordToCreateUser string
		mockUSerService.CreateUserFunc = func(user model.User, password string) (model.User, error) {
			calls += 1
			userToCreate = user
			passwordToCreateUser = password
			return user, nil
		}

		req := httptest.NewRequest(echo.POST, "/sign-up", strings.NewReader(`{
    "id": 1,
	"password":"password",
    "name": "Leanne Graham",
    "username": "Bret",
    "email": "Sincere@april.biz",
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
  }`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		srv.Echo.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, 1, calls)
		assert.Equal(t, "Sincere@april.biz", userToCreate.Email)
		assert.Equal(t, "Bret", userToCreate.Username)
		assert.Equal(t, "password", passwordToCreateUser)

	})

}
