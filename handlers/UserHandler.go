package handlers

import (
	"net/http"
	"valorize-app/models"

	"github.com/labstack/echo/v4"
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
  publicData := models.GetUserPublicProfile(&userData)
  return c.JSON(http.StatusOK, publicData)
}

func (user *UserHandler) GetLinks(c echo.Context) error {
  username := c.Param("username")
  userData, err := models.GetUserByUsername(username, *user.server.DB)
  if err != nil {
    return c.JSON(http.StatusNotFound, map[string]string{
      "error": "could not find " + username,
    })
  }
  links, err := models.GetUserLinks(&userData, *user.server.DB)
  if err != nil {
    return c.JSON(http.StatusNotFound, map[string]string{
      "error": "could not find links",
    })
  }
  return c.JSON(http.StatusOK, map[string][]models.Link{
    "links": links,
  })
}
