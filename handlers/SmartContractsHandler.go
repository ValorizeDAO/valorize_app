package handlers

import (
	"errors"
	"net/http"
	"os"
	deployer "valorize-app/contracts_deployer"
	"valorize-app/models"
	"valorize-app/services/ethereum"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
)

type ContractHandler struct {
	server *Server
	models modelsInterface
}

type modelsInterface interface {
	GetSmartContractByKey(key string) (models.SmartContract, error)
	GetSmartContractsIndex() ([]string, error)
}

func NewContractsHandler(s *Server, m modelsInterface) *ContractHandler {
	return &ContractHandler{s, m}
}

func (smartContract *ContractHandler) GetContractBytecode(c echo.Context) error {
	contractName := c.Param("key")
	contract, err := smartContract.models.GetSmartContractByKey(contractName)
	if err != nil {
		return c.JSON(http.StatusNotFound, returnErr(err))
	}
	return c.JSON(http.StatusOK, contract)
}

func (smartContract *ContractHandler) GetContractPrice(c echo.Context) error {
	contractName := c.Param("key")
	chainId := c.QueryParam("chainId")
	if len(chainId) == 0 {
		return c.JSON(http.StatusBadRequest, returnErr(errors.New("must send chainId in the request")))
	}

	client, err := ethereum.ConnectToChain(chainId)
	if err != nil {
		return c.JSON(http.StatusNotFound, returnErr(err))
	}

	deployerInstance, err := deployer.NewDeployer(common.HexToAddress(os.Getenv("DEPLOYER")), client)
	if err != nil {
		return c.JSON(http.StatusNotFound, returnErr(err))
	}

	contractInfo, err := deployerInstance.GetContractByteCodeHash(&bind.CallOpts{}, contractName)
	if err != nil {
		return c.JSON(http.StatusNotFound, returnErr(err))
	}

	return c.JSON(http.StatusOK, contractInfo.ContractParams)
}
func (smartContract *ContractHandler) GetContractKeys(c echo.Context) error {
	contracts, err := smartContract.models.GetSmartContractsIndex()
	if err != nil {
		return c.JSON(http.StatusNotFound, returnErr(err))
	}
	return c.JSON(http.StatusOK, map[string][]string{
		"smartContractKeys": contracts,
	})
}
