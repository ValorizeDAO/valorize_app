package handlers

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"valorize-app/creatortoken"
	"valorize-app/models"
	"valorize-app/services"
	"valorize-app/services/ethereum"
	"valorize-app/services/stringsUtil"
	"valorize-app/simpletoken"
	timedmint "valorize-app/timedminttoken"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jinzhu/gorm"
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

	client, err := ethereum.ConnectToChain("1")
	instance, err := creatortoken.NewCreatorToken(common.HexToAddress(user.Token.Address), client)

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
func returnErr(err error) map[string]string {
	return map[string]string{
		"error": err.Error(),
	}
}

func (token *TokenHandler) ShowToken(c echo.Context) error {
	tokenId, err := strconv.Atoi(c.Param("id"))
	tokenData, err := models.GetTokenById(uint64(tokenId), *token.server.DB)
	if err != nil {
		return c.JSON(http.StatusNotFound, returnErr(err))
	}

	client, err := ethereum.ConnectToChain(tokenData.ChainId)
	var totalSupply *big.Int
	mintAllowance := big.NewInt(0)
	nextAllowedMint := big.NewInt(0)
	maxSupply := big.NewInt(0)
	minter := ""

	switch tokenData.TokenType {
	case "simple":
		tokenInstance, err := simpletoken.NewSimpleToken(common.HexToAddress(tokenData.Address), client)
		if err != nil {
			return c.JSON(http.StatusNotFound, returnErr(err))
		}

		totalSupply, err = tokenInstance.TotalSupply(&bind.CallOpts{})
		if err != nil {
			return c.JSON(http.StatusNotFound, returnErr(err))
		}

		maxSupply = totalSupply
	case "timed_mint":
		tokenInstance, err := timedmint.NewTimedMintToken(common.HexToAddress(tokenData.Address), client)
		if err != nil {
			return c.JSON(http.StatusNotFound, returnErr(err))
		}

		totalSupply, err = tokenInstance.TotalSupply(&bind.CallOpts{})
		if err != nil {
			return c.JSON(http.StatusNotFound, returnErr(err))
		}

		maxSupply, err = tokenInstance.SupplyCap(&bind.CallOpts{})
		if err != nil {
			return c.JSON(http.StatusNotFound, returnErr(err))
		}

		mintAllowance, err = tokenInstance.MintCap(&bind.CallOpts{})
		if err != nil {
			return c.JSON(http.StatusNotFound, returnErr(err))
		}

		nextAllowedMint, err = tokenInstance.NextAllowedMintTime(&bind.CallOpts{})
		if err != nil {
			return c.JSON(http.StatusNotFound, returnErr(err))
		}

		minterAddress, err := tokenInstance.Minter(&bind.CallOpts{})
		if err != nil {
			return c.JSON(http.StatusNotFound, returnErr(err))
		}
		minter = minterAddress.String()
	case "creator":
		tokenInstance, err := creatortoken.NewCreatorToken(common.HexToAddress(tokenData.Address), client)
		if err != nil {
			return c.JSON(http.StatusNotFound, returnErr(err))
		}

		totalSupply, err = tokenInstance.TotalSupply(&bind.CallOpts{})
		if err != nil {
			return c.JSON(http.StatusNotFound, returnErr(err))
		}

	}

	return c.JSON(http.StatusOK, map[string]string{
		"name":              tokenData.Name,
		"symbol":            tokenData.Symbol,
		"address":           tokenData.Address,
		"ownerAddress":      tokenData.OwnerAddress,
		"vaultAddress":      tokenData.VaultAddress,
		"tokenType":         tokenData.TokenType,
		"contractVersion":   tokenData.ContractVersion,
		"chainId":           tokenData.ChainId,
		"totalSupply":       totalSupply.String(),
		"maxSupply":         maxSupply.String(),
		"nextMintAllowance": mintAllowance.String(),
		"nextAllowedMint":   nextAllowedMint.String(),
		"minter":            minter,
	})
}

type WalletInfo struct {
	Address string `json:"address"`
	UserId  uint   `json:"user"`
}

func (token *TokenHandler) ShowTokenAdmins(c echo.Context) error {
	tokenId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}
	var t models.Token
	if err := token.server.DB.Preload("AdminAddresses").Where("id=?", tokenId).First(&t).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
		}
	}
	var wallets []WalletInfo
	for _, s := range t.AdminAddresses {
		wallets = append(wallets, WalletInfo{s.Address, s.ID})
	}
	return c.JSON(http.StatusOK, map[string][]WalletInfo{
		"administrators": wallets,
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

	client, err := ethereum.ConnectToChain("1")
	tokenInstance, err := creatortoken.NewCreatorToken(common.HexToAddress(user.Token.Address), client)
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

	client, err := ethereum.ConnectToChain("1")
	instance, err := creatortoken.NewCreatorToken(common.HexToAddress(user.Token.Address), client)
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

type airdropRaw struct {
	Payload 	[][]string  `json:"payload"`
	MerkleRoot  string 		`json:"merkleRoot"`
}

var a airdropRaw

func (token *TokenHandler) NewAirdrop(c echo.Context) error {
	tokenId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	if err := c.Bind(&a); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	tokenData, err := models.GetTokenById(uint64(tokenId), *token.server.DB)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}
	if tokenData.TokenType == "simple" || tokenData.TokenType == "timed_mint" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "token type is not airdropable",
		})
	}
	
	rawAirdropData, err := stringsUtil.ValidateAndStringifyMap(a.Payload)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}
	airdropstruct := models.Airdrop{
		TokenID:    uint(tokenId),
		MerkleRoot: a.MerkleRoot,
		RawData:    rawAirdropData,
	}
	airdrop, err := models.NewAirdrop(*token.server.DB, airdropstruct)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	if models.NewAirdropClaim(*token.server.DB, a.Payload, airdrop.ID) != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"status":     "ok",
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

	client, err := ethereum.ConnectToChain("1")
	instance, err := creatortoken.NewCreatorToken(common.HexToAddress(creatorToken.Address), client)
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
