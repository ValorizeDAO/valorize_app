package handlers

import (
  "fmt"
  "github.com/labstack/echo/v4"
  "io"
  "net/http"
  "os"
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
      "error": "could not find " + userData.Username,
    })
  }
  publicData := map[string]string{
    "username": userData.Username,
    "name":     userData.Name,
    "id":       strconv.Itoa(int(userData.ID)),
  }
  return c.JSON(http.StatusOK, publicData)
}

func (user *UserHandler) Update(c echo.Context) error {
  username := c.Param("username")
  userData, err := models.GetUserByUsername(username, *user.server.DB)
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

  fmt.Println(file.Header)
  path := "dist/" + strconv.FormatUint(uint64(userData.ID), 10) + "_avatar.jpg"
  dst, err := os.Create(path)

  if err != nil {
    return err
  }

  defer dst.Close()

  if _, err = io.Copy(dst, src); err != nil {
    return err
  }

  return c.String(http.StatusOK, "ok then")
}
