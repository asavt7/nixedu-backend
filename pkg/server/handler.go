package server

import (
	"github.com/asavt7/nixEducation/pkg/service"
	"github.com/go-playground/validator/v10"
)

// APIHandler handler
type APIHandler struct {
	service *service.Service

	validator *validator.Validate
}

// NewAPIHandler constructs APIHandler
func NewAPIHandler(service *service.Service) *APIHandler {
	return &APIHandler{
		service:   service,
		validator: validator.New(),
	}
}
