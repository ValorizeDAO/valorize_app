package handlers

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"valorize-app/services"
	"valorize-app/services/ethereum"

	"github.com/labstack/echo/v4"
)

type EthHandler struct {
	server *Server
}

func NewEthHandler(s *Server) *EthHandler {
	return &EthHandler{s}
}

func (eth *EthHandler) Ping(c echo.Context) error {
	connection, err := eth.server.BlockChain.NetworkID(context.Background())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Could not connect to Ethereum Blockchain",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"data": "connected to " + connection.String(),
	})
}

func (eth *EthHandler) CreateWalletFromRequest(c echo.Context) error {
	password := c.FormValue("password")
	if password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "password required to generate wallet",
		})
	}
	user, _ := services.AuthUser(c, *eth.server.DB)
	address, err := ethereum.StoreUserKeystore(password, user.ID, eth.server.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"status":  "ok",
		"address": address,
	})
}

func (eth *EthHandler) DeployCreatorToken(c echo.Context) error {
	tokenName := c.FormValue("tokenName")
	tokenTicker := c.FormValue("tokenTicker")
	client, err := ethereum.Connect()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	addr, tx, _, err := ethereum.LaunchContract(client, tokenName, tokenTicker)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	user, _ := services.AuthUser(c, *eth.server.DB)

	return c.JSON(http.StatusOK, map[string]string{
		"user_id":          strconv.Itoa(int(user.ID)),
		"contract_version": "v0.1.2",
		"name":             tokenName,
		"symbol":           tokenTicker,
		"network":          os.Getenv("ETH_TESTNET"),
		"address":          addr.String(),
		"tx":               tx.Hash().String(),
	})
}
