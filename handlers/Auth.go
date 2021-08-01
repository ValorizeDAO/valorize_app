package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
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

  userStruct, err := json.Marshal(models.GetUserProfile(user))
  if err != nil {
    return c.JSON(http.StatusInternalServerError, map[string]string{
      "error": "could not find logged in user information",
    })
  }
  return c.JSON(http.StatusOK, json.RawMessage(userStruct))

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

  userStruct, err := json.Marshal(models.GetUserProfile(user))
  if err != nil {
    return c.JSON(http.StatusInternalServerError, map[string]string{
      "error": "could not find logged in user information",
    })
  }
  return c.JSON(http.StatusCreated, json.RawMessage(userStruct))

}

func (auth *AuthHandler) Show(c echo.Context) error {
  username := c.Param("username")
  user, err := models.GetUserByUsername(username, *auth.server.DB)
  if err != nil {
    return c.JSON(http.StatusNotFound, map[string]string{
      "error": "could not find " + user.Username,
    })
  }
  userStruct, err := json.Marshal(models.GetUserProfile(*user))
  if err != nil {
    return c.JSON(http.StatusInternalServerError, map[string]string{
      "error": "could not find logged in user information",
    })
  }
  return c.JSON(http.StatusOK, json.RawMessage(userStruct))

}
func (auth *AuthHandler) ShowUser(c echo.Context) error {
  user, err := services.AuthUser(c, *auth.server.DB)

  if err != nil {
    return c.JSON(http.StatusNotFound, map[string]string{
      "error": "could not find logged in user information",
    })
  }

  userStruct, err := json.Marshal(models.GetUserProfile(user))
  if err != nil {
    return c.JSON(http.StatusInternalServerError, map[string]string{
      "error": "could not find logged in user information",
    })
  }
  return c.JSON(http.StatusOK, json.RawMessage(userStruct))
}

func (auth *AuthHandler) UpdatePicture(c echo.Context) error {
  userData, err := services.AuthUser(c, *auth.server.DB)
  if err != nil {
    return c.JSON(http.StatusNotFound, map[string]string{
      "error": "could not find " + userData.Username,
    })
  }

  file, err := c.FormFile("picture")
  if err != nil {
    return c.JSON(http.StatusNotFound, map[string]string{
      "error": "could not process file",
    })
  }
  src, err := file.Open()
  if err != nil {
    return err
  }
  defer src.Close()

  filename := strconv.FormatUint(uint64(userData.ID), 10) + "_avatar.jpg"
  path := "dist/images/" + filename
  dst, err := os.Create(path)

  if err != nil {
    return err
  }

  defer dst.Close()

  if _, err = io.Copy(dst, src); err != nil {
    return err
  }

  userData.Avatar = filename
  auth.server.DB.Save(userData)

  return c.JSON(http.StatusOK, map[string]string{
    "image": filename,
  })
}

func (auth *AuthHandler) UpdateProfile(c echo.Context) error {
  userData, err := services.AuthUser(c, *auth.server.DB)
  
  if err != nil {
    return c.JSON(http.StatusNotFound, map[string]string{
      "error": "could not find logged in user",
    })
  }

  name  := c.FormValue("name")
  about := c.FormValue("about")

  userData.Name = name
  userData.About = about

  err = auth.server.DB.Save(&userData).Error
  if err != nil {
    return c.JSON(http.StatusInternalServerError, map[string]string{
      "error": "database error",
    })
  }
  userStruct, err := json.Marshal(models.GetUserProfile(userData))
  if err != nil {
    return c.JSON(http.StatusInternalServerError, map[string]string{
      "error": "could not parse user data",
    })
  }
  return c.JSON(http.StatusOK, json.RawMessage(userStruct))
}
