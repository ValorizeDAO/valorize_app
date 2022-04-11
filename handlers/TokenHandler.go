package handlers

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"valorize-app/creatortoken"
	"valorize-app/models"
	"valorize-app/services"
	"valorize-app/services/ethereum"
	"valorize-app/services/merkletree"
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

func (token *TokenHandler) Index(c echo.Context) error {
	userData, err := services.AuthUser(c, *token.server.DB)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}
	var userTokens []models.Token
	token.server.DB.Where("user_id = ?", userData.ID).Table("tokens").Find(&userTokens)

	return c.JSON(http.StatusOK, UserTokenIndexResponse{userTokens})
}

type UserTokenIndexResponse struct {
	Token []models.Token `json:"tokens"`
}
type ClaimAmountResponse struct {
	MerkleRoot  string   `json:"merkleRoot"`
	Claim       string   `json:"claim"`
	MerkleProof []string `json:"merkleProof"`
}

func (token *TokenHandler) AirdropClaimAmount(c echo.Context) error {
	tokenId, err := strconv.Atoi(c.Param("id"))
	airdropData, err := models.GetAirdropByTokenId(uint64(tokenId), *token.server.DB)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "no airdrop available for the given tokenId",
		})
	}
	wallet_address := c.Param("address")
	airdropClaimData, err := models.GetClaimAmountByAirdropID(uint64(airdropData.ID), wallet_address, *token.server.DB)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "this airdrop is not available for this address",
		})
	}
	airdropClaims, err := models.GetAllAirdropClaims(*token.server.DB, int(airdropClaimData.AirdropID))
	if err != nil {
		fmt.Println("ERROR GETTING AIRDROP DATA")
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "could not get all airdrop information",
		})
	}

	var airdropClaimsMap [][]string
	for _, claim := range airdropClaims {
		claimPair := []string{claim.WalletAddress, claim.ClaimAmount}
		airdropClaimsMap = append(airdropClaimsMap, claimPair)
	}

	rawAirdropData, err := stringsUtil.ValidateAndStringifyMap(airdropClaimsMap)
	if err != nil {
		fmt.Println("ERROR VALIDATING AIRDROP DATA")
		return c.JSON(http.StatusNotFound, returnErr(err))
	}
	proof, err := merkletree.GetMerkleProof(rawAirdropData, "[\""+wallet_address+"\",\""+airdropClaimData.ClaimAmount+"\"]")
	if err != nil {
		fmt.Println("ERROR GETTING MERKLE PROOF")
		return c.JSON(http.StatusInternalServerError, returnErr(err))
	}

	response, err := json.Marshal(ClaimAmountResponse{
		Claim:       airdropClaimData.ClaimAmount,
		MerkleRoot:  proof.Root,
		MerkleProof: proof.MerkleProof,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, returnErr(err))
	}
	fmt.Printf(string(response))
	return c.JSONBlob(http.StatusOK, response)
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

	client, _ := ethereum.ConnectToChain(user.Token.ChainId)
	instance, _ := creatortoken.NewCreatorToken(common.HexToAddress(user.Token.Address), client)

	ownerTokenBalance, _ := instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(user.Token.OwnerAddress))
	totalMinted, _ := instance.TotalSupply(&bind.CallOpts{})
	etherStaked, _ := client.BalanceAt(context.Background(), common.HexToAddress(user.Token.Address), nil)

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

type AirdropInfo struct {
	Root            [32]byte
	ClaimPeriodEnds *big.Int
	IsComplete      bool
}

type AirdropInfoResponse struct {
	AirdropIndex    *big.Int `json:"airdropOnChainIndex"`
	RootHash        string   `json:"rootHash"`
	ClaimPeriodEnds *big.Int `json:"claimPeriodEnds"`
	IsComplete      bool     `json:"isComplete"`
}

type TokenApiResponse struct {
	Name              string              `json:"name"`
	Symbol            string              `json:"symbol"`
	Address           string              `json:"address"`
	OwnerAddress      string              `json:"ownerAddress"`
	VaultAddress      string              `json:"vaultAddress"`
	TokenType         string              `json:"tokenType"`
	ContractVersion   string              `json:"contractVersion"`
	ChainId           string              `json:"chainId"`
	TotalSupply       string              `json:"totalSupply"`
	MaxSupply         string              `json:"maxSupply"`
	NextMintAllowance string              `json:"nextMintAllowance"`
	NextAllowedMint   string              `json:"nextAllowedMint"`
	Minter            string              `json:"minter"`
	AirdropSupply     string              `json:"airdropSupply"`
	Airdrop           AirdropInfoResponse `json:"airdropInfo"`
}

func (token *TokenHandler) ShowToken(c echo.Context) error {
	tokenId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, returnErr(err))
	}
	tokenData, err := models.GetTokenById(uint64(tokenId), *token.server.DB)
	if err != nil {
		return c.JSON(http.StatusNotFound, returnErr(err))
	}

	client, err := ethereum.ConnectToChain(tokenData.ChainId)
	if err != nil {
		return c.JSON(http.StatusNotFound, returnErr(err))
	}

	var totalSupply *big.Int
	mintAllowance := big.NewInt(0)
	nextAllowedMint := big.NewInt(0)
	airdropSupply := big.NewInt(0)
	maxSupply := big.NewInt(0)
	airdropInfo := AirdropInfo{}
	airdropResponse := AirdropInfoResponse{}
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

		airdropSupply, err = tokenInstance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(tokenData.Address))
		if err != nil {
			return c.JSON(http.StatusNotFound, returnErr(err))
		}

		airdropIndex, err := tokenInstance.NumberOfAirdrops(&bind.CallOpts{})
		if err != nil {
			return c.JSON(http.StatusNotFound, returnErr(err))
		}
		airdropInfo, err = tokenInstance.GetAirdropInfo(&bind.CallOpts{}, airdropIndex.Sub(airdropIndex, big.NewInt(1)))
		if err != nil {
			return c.JSON(http.StatusNotFound, returnErr(err))
		}

		airdropResponse.AirdropIndex = airdropIndex
		airdropResponse.ClaimPeriodEnds = airdropInfo.ClaimPeriodEnds
		airdropResponse.IsComplete = airdropInfo.IsComplete
		airdropResponse.RootHash = hex.EncodeToString(airdropInfo.Root[:])
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

		airdropSupply, err = tokenInstance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(tokenData.Address))
		if err != nil {
			return c.JSON(http.StatusNotFound, returnErr(err))
		}

		airdropIndex, err := tokenInstance.NumberOfAirdrops(&bind.CallOpts{})
		if err != nil {
			return c.JSON(http.StatusNotFound, returnErr(err))
		}
		airdropInfo, err = tokenInstance.GetAirdropInfo(&bind.CallOpts{}, airdropIndex.Sub(airdropIndex, big.NewInt(1)))
		if err != nil {
			return c.JSON(http.StatusNotFound, returnErr(err))
		}
		airdropResponse.AirdropIndex = airdropIndex
		airdropResponse.ClaimPeriodEnds = airdropInfo.ClaimPeriodEnds
		airdropResponse.IsComplete = airdropInfo.IsComplete
		airdropResponse.RootHash = hex.EncodeToString(airdropInfo.Root[:])
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

	return c.JSON(http.StatusOK, TokenApiResponse{
		Name:              tokenData.Name,
		Symbol:            tokenData.Symbol,
		Address:           tokenData.Address,
		OwnerAddress:      tokenData.OwnerAddress,
		VaultAddress:      tokenData.VaultAddress,
		TokenType:         tokenData.TokenType,
		ContractVersion:   tokenData.ContractVersion,
		ChainId:           tokenData.ChainId,
		TotalSupply:       totalSupply.String(),
		MaxSupply:         maxSupply.String(),
		NextMintAllowance: mintAllowance.String(),
		NextAllowedMint:   nextAllowedMint.String(),
		Minter:            minter,
		AirdropSupply:     airdropSupply.String(),
		Airdrop:           airdropResponse,
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
	if err != nil {
		return c.JSON(http.StatusNotFound, returnErr(err))
	}
	tokenInstance, err := creatortoken.NewCreatorToken(common.HexToAddress(user.Token.Address), client)
	if err != nil {
		return c.JSON(http.StatusNotFound, returnErr(err))
	}
	ethToCheckBig, ok := new(big.Int).SetString(etherToCheck, 10)

	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "could not parse ether to check",
		})
	}
	AmountForSender, AmountForOwner, _ := tokenInstance.CalculateTokenBuyReturns(&bind.CallOpts{}, ethToCheckBig)

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
	if err != nil {
		return c.JSON(http.StatusNotFound, returnErr(err))
	}

	instance, err := creatortoken.NewCreatorToken(common.HexToAddress(user.Token.Address), client)
	if err != nil {
		return c.JSON(http.StatusNotFound, returnErr(err))
	}

	tokenToCheckBig, ok := new(big.Int).SetString(tokensToCheck, 10)
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
	Payload [][]string `json:"payload"`
}

var a airdropRaw

func (token *TokenHandler) NewAirdrop(c echo.Context) error {
	tokenId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, returnErr(err))
	}

	if err := c.Bind(&a); err != nil {
		return c.JSON(http.StatusInternalServerError, returnErr(err))
	}

	tokenData, err := models.GetTokenById(uint64(tokenId), *token.server.DB)
	if err != nil {
		return c.JSON(http.StatusNotFound, returnErr(err))
	}

	if !(tokenData.TokenType == "simple" || tokenData.TokenType == "timed_mint") {
		return c.JSON(http.StatusBadRequest, returnErr(errors.New("token type is not airdropable")))
	}

	rawAirdropData, err := stringsUtil.ValidateAndStringifyMap(a.Payload)
	if err != nil {
		return c.JSON(http.StatusNotFound, returnErr(err))
	}

	merkleRoot, err := merkletree.GetMerkleRoot(rawAirdropData)
	if err != nil {
		return c.JSON(http.StatusNotFound, returnErr(err))
	}

	client, err := ethereum.ConnectToChain(tokenData.ChainId)
	if err != nil {
		return c.JSON(http.StatusNotFound, returnErr(err))
	}

	instance, err := simpletoken.NewSimpleToken(common.HexToAddress(tokenData.Address), client)
	if err != nil {
		return c.JSON(http.StatusNotFound, returnErr(err))
	}

	airdropOnChainIndex, err := instance.NumberOfAirdrops(&bind.CallOpts{})
	if err != nil {
		return c.JSON(http.StatusNotFound, returnErr(err))
	}

	onchainIndex, err := strconv.Atoi(airdropOnChainIndex.String())
	if err != nil {
		return c.JSON(http.StatusNotFound, returnErr(errors.New("Error getting airdrop index from token contract")))
	}

	airdropstruct := models.Airdrop{
		TokenID:      uint(tokenId),
		OnChainIndex: uint(onchainIndex),
		MerkleRoot:   merkleRoot,
	}
	airdrop, err := models.NewAirdrop(*token.server.DB, airdropstruct)
	if err != nil {
		return c.JSON(http.StatusNotFound, returnErr(err))
	}

	if models.NewAirdropClaim(*token.server.DB, a.Payload, airdrop.ID) != nil {
		return c.JSON(http.StatusInternalServerError, returnErr(err))
	}

	return c.JSON(http.StatusOK, map[string]string{
		"status":     "ok",
		"merkleRoot": merkleRoot,
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
