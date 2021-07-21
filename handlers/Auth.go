package handlers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
	models "valorize-app/models"
	"valorize-app/services/ethereum"
)

type AuthHandler struct {
	DB *gorm.DB
}

func (auth *AuthHandler) Login(c echo.Context) error {

	username := c.FormValue("username")
	password := c.FormValue("password")

	user := models.User{}

	err := auth.DB.First(&user, "username = ?", username).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": username + " does not exist",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "incorrect password for " + username,
		})

	}

	_user := user.Username
	_expiration := time.Now().Add(time.Hour * 144).Unix()
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = _user
	claims["exp"] = _expiration

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	cookie := tokenCookie(t)
	c.SetCookie(cookie)

	return c.JSON(http.StatusCreated, map[string]string{
		"token":   t,
		"user":    _user,
		"expires": strconv.Itoa(int(_expiration)),
	})
}

func (auth *AuthHandler) Register(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	user := models.User{
		Email:    email,
		Username: username,
		Password: string(hash),
	}

	if auth.DB.Create(&user).Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Could not create user",
		})
	}

	address, err := ethereum.NewKeystore(password)
	wallet := models.Wallet{
		Address: address,
		UserId:  user.ID,
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error creating ethereum wallet",
		})
	}
	if auth.DB.Create(&wallet).Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Could not store wallet",
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"email": email,
		"user":  username,
	})
}

func tokenCookie(t string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = t
	cookie.Expires = time.Now().Add(144 * time.Hour)
	cookie.HttpOnly = true
	return cookie
}
