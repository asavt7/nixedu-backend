package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags healthcheck
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func healthCheck(context echo.Context) error {
	return context.JSON(http.StatusOK, map[string]interface{}{})
}
