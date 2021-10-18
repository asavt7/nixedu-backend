package server

import (
	"github.com/asavt7/nixEducation/pkg/service"
	"github.com/go-playground/validator/v10"
)

type ApiHandler struct {
	service *service.Service

	validator *validator.Validate
}

func NewApiHandler(service *service.Service) *ApiHandler {
	return &ApiHandler{
		service:   service,
		validator: validator.New(),
	}
}
