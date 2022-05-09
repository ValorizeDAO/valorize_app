package handlers_test

import (
	"errors"
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
	tests := []getContractByteCodeTest{
		{
			name:              "Should return contract bytecode",
			shouldReturnError: false,
			expectedStatus:    http.StatusOK,
			expectedBody:      "{\"id\":0,\"key\":\"test_key\",\"byte_code\":\"0x12a34\"}\n",
		},
		{
			name:              "Should return error",
			shouldReturnError: true,
			expectedStatus:    http.StatusNotFound,
			expectedBody:      "{\"error\":\"error getting smart contract\"}\n",
		},
	}
	for _, test := range tests {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/:key")
		c.SetParamNames("key")
		c.SetParamValues("test_key")

		contracts := handlers.NewContractsHandler(&handlers.Server{}, mockModels{test.shouldReturnError})
		if assert.NoError(t, contracts.GetContractBytecode(c)) {
			assert.Equal(t, test.expectedStatus, rec.Code)
			assert.Equal(t, test.expectedBody, rec.Body.String())
		}
	}
}

type getContractByteCodeTest struct {
	name              string
	shouldReturnError bool
	expectedStatus    int
	expectedBody      string
}

type mockModels struct {
	ShouldReturnError bool
}

func (m mockModels) GetSmartContractByKey(key string) (models.SmartContract, error) {
	if m.ShouldReturnError {
		return models.SmartContract{}, errors.New("error getting smart contract")
	}
	return models.SmartContract{Key: key, ByteCode: "0x12a34"}, nil
}
