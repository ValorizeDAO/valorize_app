package handlers

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"valorize-app/contracts"
	"valorize-app/models"
	"valorize-app/services"
	"valorize-app/services/ethereum"
	"valorize-app/simpletoken"

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

type PublicTokenResponse struct {
	Token     *models.Token     `json:"token"`
	PriceData map[string]string `json:"price_data"`
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
	etherStaked, err := client.BalanceAt(context.Background(), common.HexToAddress(user.Token.Address), nil)

	return c.JSON(http.StatusOK, PublicTokenResponse{
		Token: &user.Token,
		PriceData: map[string]string{
			"owner_balance": ownerTokenBalance.String(),
			"total_minted":  totalMinted.String(),
			"ether_staked":  etherStaked.String(),
		},
	})
}

func (token *TokenHandler) ShowToken(c echo.Context) error {
	tokenId, err := strconv.Atoi(c.Param("id"))
	tokenData, err := models.GetTokenById(uint64(tokenId), *token.server.DB)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	client, err := ethereum.ConnectToChain(tokenData.ChainId)
	instance, err := simpletoken.NewSimpleToken(common.HexToAddress(tokenData.Address), client)

	totalSupply, err := instance.TotalSupply(&bind.CallOpts{})

	return c.JSON(http.StatusOK, map[string]string{
		"name":            tokenData.Name,
		"symbol":          tokenData.Symbol,
		"address":         tokenData.Address,
		"ownerAddress":    tokenData.OwnerAddress,
		"vaultAddress":    tokenData.VaultAddress,
		"tokenType":       tokenData.TokenType,
		"contractVersion": tokenData.ContractVersion,
		"chainId":         tokenData.ChainId,
		"totalSupply":     totalSupply.String(),
	})
}

func (token *TokenHandler) GetTokenStakingRewards(c echo.Context) error {
	username := c.Param("username")
	etherToCheck := c.FormValue("etherToCheck")
	if etherToCheck == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "etherToCheck param required",
		})
	}
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
	tokenInstance, err := contracts.NewCreatorToken(common.HexToAddress(user.Token.Address), client)
	ethToCheckBig, ok := new(big.Int).SetString(etherToCheck, 10)

	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "could not parse ether to check",
		})
	}
	AmountForSender, AmountForOwner, err := tokenInstance.CalculateTokenBuyReturns(&bind.CallOpts{}, ethToCheckBig)

	return c.JSON(http.StatusOK, map[string]string{
		"toBuyer": AmountForSender.String(),
		"toOwner": AmountForOwner.String(),
	})
}

func (token *TokenHandler) GetTokenSellingRewards(c echo.Context) error {
	username := c.Param("username")
	tokensToCheck := c.FormValue("tokenToCheck")
	fmt.Printf("tokensToCheck: %s", tokensToCheck)
	if tokensToCheck == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "tokenToCheck param required",
		})
	}
	user, err := models.GetUserByUsername(username, *token.server.DB)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "could not find " + username,
		})
	}

	if !user.HasDeployedToken {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": user.Username + " has not deployed a token",
		})
	}

	client, err := ethereum.MainnetConnection()
	instance, err := contracts.NewCreatorToken(common.HexToAddress(user.Token.Address), client)
	tokenToCheckBig, ok := new(big.Int).SetString(tokensToCheck, 10)
	fmt.Printf("tokensToCheckBig: %s", tokenToCheckBig)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "could not parse tokens to check",
		})
	}
	saleReturnAmount, err := instance.CalculateTotalSaleReturn(&bind.CallOpts{}, tokenToCheckBig)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "could not calculate sale return: " + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"total_eth": saleReturnAmount.Text(10),
	})
}

type TokenBalanceResponse struct {
	TotalBalance string          `json:"total_balance"`
	Wallets      []WalletBalance `json:"wallets"`
}

type WalletBalance struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
}

func (token *TokenHandler) GetGasPriceToLaunchToken(c echo.Context) error {
	gasPrice, err := ethereum.GetGasPrice("MAINNET")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "could not get gas price, " + err.Error(),
		})
	}

	currentGasPriceToLaunchContract := gasPrice * int64(3000000)
	return c.JSON(http.StatusOK, map[string]string{
		"gas_price": strconv.FormatUint(uint64(currentGasPriceToLaunchContract), 10),
	})
}

func (token *TokenHandler) GetCoinBalanceForAuthUser(c echo.Context) error {
	user, _ := services.AuthUser(c, *token.server.DB)

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

	var WalletBalances []WalletBalance
	totalBalance := new(big.Int)
	for _, wallet := range wallets {
		balance, err := instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(wallet))
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "error getting balance from contract for address " + wallet,
			})
		}
		if balance.Cmp(big.NewInt(0)) > 0 {
			WalletBalances = append(WalletBalances, WalletBalance{wallet, balance.String()})
			totalBalance.Add(totalBalance, balance)
		}
	}

	return c.JSON(http.StatusOK, TokenBalanceResponse{
		TotalBalance: totalBalance.String(),
		Wallets:      WalletBalances,
	})
}
