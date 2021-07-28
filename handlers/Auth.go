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
	"valorize-app/services/ethereum"
)

type AuthHandler struct {
	server *Server
}

func NewAuthHandler(s *Server) *AuthHandler {
	return &AuthHandler{
		server: s,
	}
}

func (auth *AuthHandler) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	user := models.User{}

	err := auth.server.DB.First(&user, "username = ?", username).Error
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

	if auth.server.DB.Create(&user).Error != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": user.Username + " already exists",
		})
	}

	go ethereum.StoreUserKeystore(password, user.ID, auth.server.DB)

	return c.JSON(http.StatusCreated, map[string]string{
		"email": email,
		"user":  username,
	})
}

func (auth *AuthHandler) Show(c echo.Context) error {
	username := c.Param("username")
	user, err := models.GetUserByUsername(username, *auth.server.DB)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "could not find " + user.Username,
		})
	}
	publicData := map[string]string{
		"username": user.Username,
		"name":     user.Name,
		"id":       strconv.Itoa(int(user.ID)),
	}
	return c.JSON(http.StatusOK, publicData)
}
