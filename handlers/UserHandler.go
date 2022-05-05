package handlers

import (
	"net/http"
	"valorize-app/models"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	server *Server
	models *models.Model
}

func NewUserHandler(s *Server, m *models.Model) *UserHandler {
	return &UserHandler{s, m}
}
func (user *UserHandler) Show(c echo.Context) error {
	username := c.Param("username")
	userData, err := user.models.GetUserByUsername(username)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "could not find " + username,
		})
	}
	publicData := user.models.GetUserPublicProfile(&userData)
	return c.JSON(http.StatusOK, publicData)
}

func (user *UserHandler) GetLinks(c echo.Context) error {
	username := c.Param("username")
	userData, err := user.models.GetUserByUsername(username)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "could not find " + username,
		})
	}
	links, err := user.models.GetUserLinks(&userData)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "could not find links",
		})
	}
	return c.JSON(http.StatusOK, map[string][]models.Link{
		"links": links,
	})
}
