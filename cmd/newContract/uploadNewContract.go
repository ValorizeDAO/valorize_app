package main

import (
	"fmt"
	"os"
	"valorize-app/config"
	"valorize-app/db"
	"valorize-app/models"

	"github.com/jinzhu/gorm"
)

func main() {
	cfg := config.NewConfig()
	database := db.Init(cfg)
	defer database.Close()
	contrackKey := os.Args[1]
	bytecode := os.Args[2]
	c, err := AddContractByteCode(database, contrackKey, bytecode)
	if err == nil {
		fmt.Printf("Added new contract [%v] %v", c.ID, c.Key)
	}

}

func AddContractByteCode(db *gorm.DB, key string, byteCodeHex string) (models.SmartContract, error) {
	var contractInfo models.SmartContract
	err := db.Where(
		models.SmartContract{Key: key},
	).Attrs(
		models.SmartContract{ByteCode: byteCodeHex},
	).FirstOrCreate(&contractInfo).Error
	if err != nil {
		return contractInfo, err
	}
	return contractInfo, err
}
