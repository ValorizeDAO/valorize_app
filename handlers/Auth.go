package handlers

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	models "valorize-app/models"
	"valorize-app/services"
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

	token, expiration, err := services.NewToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "could not create token",
		})
	}

	cookie := services.CreateTokenCookie(token)
	c.SetCookie(cookie)

	return c.JSON(http.StatusCreated, map[string]string{
		"token":   token,
		"user":    user.Username,
		"expires": strconv.FormatInt(expiration, 10),
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

	return c.JSON(http.StatusCreated, map[string]string{
		"email": email,
		"user":  username,
	})
}
