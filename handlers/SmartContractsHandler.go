package handlers

import (
	"net/http"
	"valorize-app/models"

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
func (smartContract *ContractHandler) GetContractKeys(c echo.Context) error {
	contracts, err := smartContract.models.GetSmartContractsIndex()
	if err != nil {
		return c.JSON(http.StatusNotFound, returnErr(err))
	}
	return c.JSON(http.StatusOK, map[string][]string{
		"smartContractKeys": contracts,
	})
}
