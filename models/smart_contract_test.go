package models_test

import (
	"testing"
	"valorize-app/models"

	_ "github.com/go-sql-driver/mysql"
)

type getSmartContractByKeyTest struct {
	key              string
	expectedBytecode string
}

func TestGetSmartContractByKey(t *testing.T) {
	prepareTestDatabase()
	tests := []getSmartContractByKeyTest{
		{"simple_token_v0.1.0", "0x123"},
		{"timedMint_token_v0.1.0", "0x123a4"},
	}
	mdl := models.NewModels(gormDb)
	for _, test := range tests {
		output, err := mdl.GetSmartContractByKey(test.key)
		if err != nil {
			t.Error(err)
		}
		if output.ByteCode != test.expectedBytecode {
			t.Errorf("\nGetSmartContractByKey(%v) returns bytecode: returned %v, want %v",
				test.key, output.ByteCode, test.expectedBytecode)
		}
	}
}
