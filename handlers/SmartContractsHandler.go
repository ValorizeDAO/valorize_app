package handlers

import (
	"valorize-app/models"

	"github.com/labstack/echo/v4"
)

type ContractHandler struct {
	server *Server
	models *models.Model
}

func NewContractsHandler(s *Server, m *models.Model) *ContractHandler {
	return &ContractHandler{s, m}
}

func (smartContract *ContractHandler) GetContractBytecode(c echo.Context) error {
	smartContract.models.GetSmartContractByKey("contract_bytecode")
	return nil
}
