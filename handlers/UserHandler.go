package handlers

import (
  "github.com/labstack/echo/v4"
  "net/http"
  "strconv"
  "valorize-app/models"
)

type UserHandler struct {
  server *Server
}

func NewUserHandler(s *Server) *UserHandler {
  return &UserHandler{
    server: s,
  }
}
func (user *UserHandler) Show(c echo.Context) error {
  username := c.Param("username")
  userData, err := models.GetUserByUsername(username, *user.server.DB)
  if err != nil {
    return c.JSON(http.StatusNotFound, map[string]string{
      "error": "could not find " + username,
    })
  }
  publicData := map[string]string{
    "username": userData.Username,
    "name":     userData.Name,
    "id":       strconv.Itoa(int(userData.ID)),
    "avatar":   userData.Avatar,
    "about":    userData.About,
  }
  return c.JSON(http.StatusOK, publicData)
}

