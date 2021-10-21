package service

import (
	"errors"
	"fmt"
	mock_storage "github.com/asavt7/nixedu/backend/mocks/pkg/storage"
	"github.com/asavt7/nixedu/backend/pkg/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type UserMatcher struct {
	User     model.User
	Password string
}

func (u UserMatcher) Matches(x interface{}) bool {
	switch x.(type) {
	case model.User:
		return x.(model.User).Username == u.User.Username && x.(model.User).Email == u.User.Email && checkPassword(x.(model.User).PasswordHash, u.Password) == nil
	default:
		return false
	}
}

func (u UserMatcher) String() string {
	return fmt.Sprintf("username=%v password=%s", u.User, u.Password)
}

func TestUserServiceImpl_CreateUser(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	userStorage := mock_storage.NewMockUserStorage(controller)
	userService := UserServiceImpl{userStorage}

	t.Run("ok", func(t *testing.T) {
		expectedCreatedUser := model.User{Email: "email", Username: "username", ID: 1}
		userStorage.EXPECT().Create(UserMatcher{
			User:     expectedCreatedUser,
			Password: "password",
		}).Return(expectedCreatedUser, nil)

		user, err := userService.CreateUser(model.User{Email: "email", Username: "username"}, "password")
		if err != nil {
			t.Errorf("err should be nil")
		}

		assert.Equal(t, expectedCreatedUser, user)
	})

	t.Run("error", func(t *testing.T) {
		expectedCreatedUser := model.User{Email: "email", Username: "username", ID: 1}
		userStorage.EXPECT().Create(UserMatcher{
			User:     expectedCreatedUser,
			Password: "password",
		}).Return(expectedCreatedUser, errors.New("cannot create user"))

		_, err := userService.CreateUser(model.User{Email: "email", Username: "username"}, "password")
		if err == nil {
			t.Errorf("err should not nil")
		}
	})
}

func TestUserServiceImpl_GetUserById(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	userStorage := mock_storage.NewMockUserStorage(controller)
	userService := UserServiceImpl{userStorage}

	userID := 1

	t.Run("ok", func(t *testing.T) {
		expectedUser := model.User{Email: "email", Username: "username", ID: userID}
		userStorage.EXPECT().GetByID(userID).Return(expectedUser, nil)

		user, err := userService.GetUserByID(userID)
		if err != nil {
			t.Errorf("err should be nil")
		}

		assert.Equal(t, expectedUser, user)
	})

	t.Run("error", func(t *testing.T) {
		expectedUser := model.User{Email: "email", Username: "username", ID: userID}
		userStorage.EXPECT().GetByID(userID).Return(expectedUser, model.UserNotFoundErr{ID: userID})

		_, err := userService.GetUserByID(userID)
		if err == nil {
			t.Errorf("err should not nil")
		}
		assert.Equal(t, model.UserNotFoundErr{ID: userID}, err)
	})
}

func TestUserServiceImpl_GetUserByEmail(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	userStorage := mock_storage.NewMockUserStorage(controller)
	userService := UserServiceImpl{userStorage}

	email := "email@google.com"

	t.Run("ok", func(t *testing.T) {
		expectedUser := model.User{Email: email, Username: "username", ID: 1}
		userStorage.EXPECT().FindByEmail(email).Return(expectedUser, nil)

		user, err := userService.GetUserByEmail(email)
		if err != nil {
			t.Errorf("err should be nil")
		}

		assert.Equal(t, expectedUser, user)
	})

	t.Run("error", func(t *testing.T) {
		expectedUser := model.User{Email: email, Username: "username", ID: 1}
		userStorage.EXPECT().FindByEmail(email).Return(expectedUser, model.UserNotFoundErr{})

		_, err := userService.GetUserByEmail(email)

		if err == nil {
			t.Errorf("err should not nil")
		}
		assert.Equal(t, model.UserNotFoundErr{}, err)
	})
}
