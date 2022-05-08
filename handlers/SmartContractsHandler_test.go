package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"valorize-app/handlers"
	"valorize-app/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetContractBytecode(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:contract_key")
	c.SetParamNames("contract_key")
	c.SetParamValues("key")

	contracts := handlers.NewContractsHandler(&handlers.Server{}, mockModels{0})
	if assert.NoError(t, contracts.GetContractBytecode(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"id\":0,\"key\":\"contract_bytecode\",\"byte_code\":\"0x12a34\"}\n", rec.Body.String())
	}
}

type mockModels struct {
	Count int
}

func (m mockModels) GetSmartContractByKey(key string) (models.SmartContract, error) {
	return models.SmartContract{Key: key, ByteCode: "0x12a34"}, nil
}
