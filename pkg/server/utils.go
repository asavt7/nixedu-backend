package server

import (
	"github.com/labstack/echo/v4"
	"strings"
)

type message struct {
	Message string `json:"message" xml:"message"`
}

func response(status int, body interface{}, c echo.Context) error {
	ctype := c.Request().Header.Get(echo.HeaderContentType)
	acceptType := c.Request().Header.Get(echo.HeaderAccept)
	if len(acceptType) == 0 {
		acceptType = ctype
	}
	switch {
	case strings.Contains(acceptType, echo.MIMEApplicationJSON):
		return c.JSON(status, body)
	case strings.Contains(acceptType, echo.MIMEApplicationXML), strings.Contains(acceptType, echo.MIMETextXML):
		return c.XML(status, body)
	default:
		return c.JSON(status, body)
	}
}
