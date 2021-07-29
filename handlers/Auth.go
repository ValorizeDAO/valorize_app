package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"valorize-app/models"
	"valorize-app/services"
	"valorize-app/services/ethereum"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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

  token, _, err := services.NewToken(user)
  if err != nil {
    return c.JSON(http.StatusInternalServerError, map[string]string{
      "error": "could not create token",
    })
  }

  cookie := services.CreateTokenCookie(token)
  c.SetCookie(cookie)

  return c.JSON(http.StatusCreated, map[string]string{
    "id":        strconv.FormatUint(uint64(user.ID), 10),
    "name":      user.Name,
    "username":  user.Username,
    "email":     user.Email,
  })
}

func (auth *AuthHandler) Logout(c echo.Context) error {
  cookie, err := c.Cookie("token")
  if err != nil {
    return c.JSON(http.StatusOK, map[string]string{
      "status": "not logged in",
    })
  }
  cookie.Value = ""
  cookie.MaxAge = -1
  c.SetCookie(cookie)

  return c.JSON(http.StatusOK, map[string]string{
    "status": "goodbye",
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
func (auth *AuthHandler) ShowUser(c echo.Context) error {
  user, err := services.AuthUser(c, *auth.server.DB)

  if err != nil {
    return c.JSON(http.StatusNotFound, map[string]string{
      "error": "could not find logged in user information",
    })
  }

  return c.JSON(http.StatusOK, map[string]string{
    "id":        strconv.FormatUint(uint64(user.ID), 10),
    "name":      user.Name,
    "username":  user.Username,
    "email":     user.Email,
  })
}
