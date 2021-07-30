package services

import (
	"errors"
	"net/http"
	"time"
	"valorize-app/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type TokenClaims struct {
	Username string `json:"username"`
	ID       uint
	jwt.StandardClaims
}

func GetTokenFromCookie(c echo.Context) (*jwt.Token, error) {
	cookie, err := c.Cookie("token")
	if err != nil {
		return nil, errors.New("no token cookie found")
	}
	claims := &TokenClaims{}
	token, err := jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
		//TODO: get better key management for app
		return []byte("secret"), nil
	})
	return token, err
}

func AuthUser(c echo.Context, DB gorm.DB) (models.User, error) {
	token, err := GetTokenFromCookie(c)
	if err != nil {
		return models.User{}, errors.New("error getting user")
	}
	user := models.User{}
	if DB.First(&user, "username = ?", token.Claims.(*TokenClaims).Username).Error != nil {
		return models.User{}, errors.New("could not find user in database")
	}
	return user, nil
}

func NewToken(user models.User) (string, int64, error) {
	_expiration := time.Now().Add(time.Hour * 144).Unix()
	claims := &TokenClaims{
		user.Username,
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: _expiration,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	return t, _expiration, err
}

func CreateTokenCookie(token string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(144 * time.Hour)
	cookie.HttpOnly = true
	cookie.SameSite = 1
	return cookie
}
