package handlers_test

import (
	"net/http"
	"testing"
	"valorize-app/handlers"
	"valorize-app/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestGetContractBytecode(t *testing.T) {
	type ModelsTest interface {
		GetSmartContractByKey(string) (models.SmartContract, error)
	}
	type mockInterface struct {
		ModelsTest
	}
	m := mockInterface{}

	contracts := handlers.NewContractsHandler(&handlers.Server{}, m)

	// Assertions
	if assert.NoError(t, h.createUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}
