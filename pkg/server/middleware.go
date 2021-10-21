package server

import (
	"github.com/asavt7/nixEducation/pkg/service"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func parseAccessToken() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:                  &service.Claims{},
		SigningKey:              []byte(service.GetJWTSecret()),
		TokenLookup:             "header: Authorization,cookie:" + accessTokenCookieName,
		ErrorHandlerWithContext: jwtErrorChecker,
		SuccessHandler: func(c echo.Context) {
			tok := c.Get("user")
			accessToken := tok.(*jwt.Token)
			claims := accessToken.Claims.(*service.Claims)
			userID := claims.UserID
			c.Set(currentUserID, userID)
		},
	})
}

func (h *APIHandler) tokenRefresherMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		tok := c.Get("user")
		if tok == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}
		accessToken := tok.(*jwt.Token)
		claims := accessToken.Claims.(*service.Claims)

		err := h.service.ValidateAccessToken(claims)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}

		if h.service.AuthorizationService.IsNeedToRefresh(claims) {
			rc, err := c.Cookie(refreshTokenCookieName)
			if err == nil && rc != nil {
				refreshClaims, err := h.service.ParseRefreshTokenToClaims(rc.Value)
				if err != nil {
					if err == jwt.ErrSignatureInvalid {
						return echo.NewHTTPError(http.StatusUnauthorized, "invalid token signature")
					}
					return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
				}

				if claims.UserID != refreshClaims.UserID {
					return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
				}

				err = h.service.AuthorizationService.ValidateRefreshToken(refreshClaims)
				if err != nil {
					return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
				}

				_, _, err = h.generateTokensAndSetCookies(claims.UserID, c)
				if err != nil {
					return echo.NewHTTPError(http.StatusUnauthorized, "Token is incorrect")
				}
			}
		}

		return next(c)
	}
}

func jwtErrorChecker(err error, _ echo.Context) error {
	return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
}
