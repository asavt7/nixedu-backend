package server

import (
	"github.com/asavt7/nixEducation/pkg/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type signUpUserInput struct {
	Password string `json:"password"`
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
	return c.JSON(http.StatusCreated, createdUser)
}
