package handlers

import (
  "github.com/labstack/echo/v4"
  "net/http"
  "valorize-app/models"
)

type WalletHandler struct {
  server *Server
}

func NewWalletHandler(s *Server) *WalletHandler {
  return &WalletHandler{
    server: s,
  }
}
func (wallet *WalletHandler) Index(c echo.Context) error {
  username := c.Param("username")
  userData, err := models.GetUserByUsername(username, *wallet.server.DB)
  if err != nil {
    return c.JSON(http.StatusNotFound, map[string]string{
      "error": "could not find " + userData.Username,
    })
  }
  var userWallets []models.Wallet
  wallet.server.DB.Where( "user_id = ?", userData.ID).Table("wallets").Select("address").Find(&userWallets)
  var addresses = make([]string, len(userWallets))
  for i, wallet := range userWallets {
    addresses[i] = wallet.Address
  }
  return c.JSON(http.StatusOK, UserWalletResponse{userData.Username, addresses})
}

type UserWalletResponse struct{
  Username string  `json:"username"`
  Wallets   []string `json:"wallet"`
}

