package handlers

import (
	"math/big"
	"net/http"
	"strconv"
	"valorize-app/contracts"
	"valorize-app/models"
	"valorize-app/services/ethereum"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
)

type TokenHandler struct {
	server *Server
}

func NewTokenHandler(s *Server) *TokenHandler {
	return &TokenHandler{
		server: s,
	}
}

func (token *TokenHandler) Show(c echo.Context) error {
	username := c.Param("username")
	user, err := models.GetUserByUsername(username, *token.server.DB)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "could not find token for " + username,
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
	totalMinted, err := instance.TotalSupply(&bind.CallOpts{})
	etherStaked, err := instance.GetEthBalance(&bind.CallOpts{})

	return c.JSON(http.StatusOK, map[string]string{
		"owner_balance": ownerTokenBalance.String(),
		"total_minted":  totalMinted.String(),
		"ether_staked":  etherStaked.String(),
	})
}


func (token *TokenHandler) GetTokenStakingRewards(c echo.Context) error {
	username := c.Param("username")
	etherToCheck := c.FormValue("etherToCheck")
	user, err := models.GetUserByUsername(username, *token.server.DB)
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
	n := new(big.Int)
	ethToCheckBig, _ := n.SetString(etherToCheck, 10)
	AmountForSender, AmountForOwner, err := instance.CalculateTokenBuyReturns(&bind.CallOpts{}, ethToCheckBig)

	return c.JSON(http.StatusOK, map[string]string{
		"toBuyer": AmountForSender.String(),
		"toOwner": AmountForOwner.String(),
	})
}


type TokenBalanceResponse struct {
	TotalBalance *big.Int `json:"total_balance"`
	Wallets []WalletBalance `json:"wallets"`
}

type WalletBalance struct {
	Address string `json:"address"`
	Balance *big.Int `json:"balance"`
}

func (token *TokenHandler) GetBalanceForCoinForUser(c echo.Context) error {
	username := c.FormValue("username")
	if username == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "username parameter is required on request",
		})
	}

	creatorTokenId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "token id invalid",
		})
	}
	creatorToken, err := models.GetTokenById(creatorTokenId, *token.server.DB)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "could not find token",
		})
	}

	user, err := models.GetUserByUsername(username, *token.server.DB)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "could not find user with username " + username,
		})
	}
	if !user.HasDeployedToken {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": user.Username + " has not deployed a token",
		})
	}

	wallets, err := models.GetAllWalletsByUserId(user.ID, *token.server.DB)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "could not find wallets for user with id " + strconv.FormatUint(uint64(user.ID), 10),
		})
	}

	client, err := ethereum.MainnetConnection()
	instance, err := contracts.NewCreatorToken(common.HexToAddress(creatorToken.Address), client)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "could not find token",
		})
	}

	var WalletBalances []WalletBalance;
	totalBalance := new(big.Int)
	for _, wallet := range wallets {
		balance, err := instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(wallet))
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "error getting balance from contract for address " + wallet,
			})
		}
		if balance.Cmp(big.NewInt(0)) > 0 {
			WalletBalances = append(WalletBalances, WalletBalance{ wallet, balance,})
			totalBalance.Add(totalBalance, balance)
		}
	}

	return c.JSON(http.StatusOK, TokenBalanceResponse{
		TotalBalance: totalBalance,
		Wallets: WalletBalances,
	})
}
