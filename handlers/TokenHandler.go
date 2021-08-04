package handlers

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"net/http"
	"valorize-app/contracts"
	"valorize-app/models"
	"valorize-app/services/ethereum"
)

type TokenHandler struct {
	server *Server
}

func NewTokenHandler(s *Server) *TokenHandler {
	return &TokenHandler{
		server: s,
	}
}

func (auth *TokenHandler) Show(c echo.Context) error {
	username := c.Param("username")
	user, err := models.GetUserByUsername(username, *auth.server.DB)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "could not find token for " + user.Username,
		})
	}

	if !user.HasDeployedToken {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": user.Username + " has not deployed a token",
		})
	}

	client, err := ethereum.MainnetConnection()
	instance, err := contracts.NewCreatorToken(common.HexToAddress(user.Token.Address), client)

	ownerTokenBalance, err := instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(user.Token.OwnerAddress))
	totalMinted, err := instance.TotalMinted(&bind.CallOpts{})
	etherStaked, err := instance.GetEthBalance(&bind.CallOpts{})

	return c.JSON(http.StatusOK, map[string]string{
		"owner_balance": ownerTokenBalance.String(),
		"total_minted":  totalMinted.String(),
		"ether_staked":  etherStaked.String(),
	})

}
