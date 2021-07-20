package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("token")
		if err != nil {
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
			//TODO: get better key management for app
			return []byte("secret"), nil
		})
		if token.Valid {
			fmt.Print("\n\n=====\nvalid\n=====\n\n")
			return next(c)
		} else {
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}
	}
}
