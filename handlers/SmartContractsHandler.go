package handlers

import (
	"github.com/labstack/echo/v4"
)

type ContractHandler struct {
	server *Server
	models modelsInterface
}

type modelsInterface interface {
	GetSmartContractByKey(key string)
}

func NewContractsHandler(s *Server, m modelsInterface) *ContractHandler {
	return &ContractHandler{s, m}
}

func (smartContract *ContractHandler) GetContractBytecode(c echo.Context) error {
	smartContract.models.GetSmartContractByKey("contract_bytecode")
	return nil
}
