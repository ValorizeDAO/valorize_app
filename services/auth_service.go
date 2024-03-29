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
	token, _, err := GetTokenClaims(cookie.Value)
	return token, err
}

func GetTokenClaims(tokenRaw string) (*jwt.Token, *TokenClaims, error) {
	claims := &TokenClaims{}
	token, err := jwt.ParseWithClaims(tokenRaw, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if tokenClaims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return token, tokenClaims, err
	} else {
		return &jwt.Token{}, &TokenClaims{}, errors.New("token invalid")
	}
}

func AuthUser(c echo.Context, DB gorm.DB) (models.User, error) {
	token, err := GetTokenFromCookie(c)
	if err != nil {
		return models.User{}, errors.New("error getting user")
	}
	model := models.NewModels(&DB)
	user, err := model.GetUserProfileByUsername(token.Claims.(*TokenClaims).Username)
	if err != nil {
		return models.User{}, errors.New("could not find user in database")
	}
	return user, nil
}

func NewToken(user models.User, duration time.Duration) (string, int64, error) {
	if duration == 0 {
		duration = 144
	}
	_expiration := time.Now().Add(time.Hour * duration).Unix()
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

func GetUserWallets(user *models.User, db gorm.DB) []models.Wallet {
	var userWallets []models.Wallet
	db.Where("user_id = ?", user.ID).Table("wallets").Find(&userWallets)
	return userWallets
}
