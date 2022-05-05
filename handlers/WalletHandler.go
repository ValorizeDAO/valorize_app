package handlers

import (
	"net/http"
	"valorize-app/models"
	"valorize-app/services"

	"github.com/labstack/echo/v4"
)

type WalletHandler struct {
	server *Server
	models *models.Model
}

func NewWalletHandler(s *Server, m *models.Model) *WalletHandler {
	return &WalletHandler{s, m}
}

func (wallet *WalletHandler) Index(c echo.Context) error {
	userData, err := services.AuthUser(c, *wallet.server.DB)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "could not find " + userData.Username,
		})
	}
	var userWallets []models.Wallet
	wallet.server.DB.Where("user_id = ?", userData.ID).Table("wallets").Select("address").Find(&userWallets)
	var addresses = make([]string, len(userWallets))
	for i, wallet := range userWallets {
		addresses[i] = wallet.Address
	}
	return c.JSON(http.StatusOK, UserWalletResponse{addresses})
}

type UserWalletResponse struct {
	Wallets []string `json:"wallet"`
}
