package server

import (
	"github.com/asavt7/nixEducation/pkg/service"
)

type ApiHandler struct {
	service *service.Service
}

func NewApiHandler(service *service.Service) *ApiHandler {
	return &ApiHandler{service: service}
}
