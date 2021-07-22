package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"valorize-app/services"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := services.GetTokenFromCookie(c)
		if err != nil {
			return c.String(http.StatusUnauthorized, err.Error())
		}
		if token.Valid {
			return next(c)
		} else {
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}
	}
}
