package main

import (
	"fmt"
	"os"
	"strconv"
	"valorize-app/config"
	"valorize-app/db"
	"valorize-app/models"
	"valorize-app/services/ethereum"

	"github.com/jinzhu/gorm"
)

func main() {
	cfg := config.NewConfig()
	database := db.Init(cfg)
	defer database.Close()
	userIndex := os.Args[1]
	tkn, err := NewToken(database, userIndex)
	if err == nil {
		fmt.Print("Token Seeded", tkn.Address)
	} 

}

func NewToken(db *gorm.DB, i string) (models.Token, error) {
	ethInstance, err := ethereum.Connect()
	address, tx, _, err := ethereum.LaunchContract(ethInstance, "NewCoin", "NU")
	if err != nil {
		fmt.Println("Error connecting to Ethereum")
	}
	userId, err := strconv.ParseUint(i, 10, 64)
	if err != nil {
		fmt.Println("Error converting userId to int")
	}
	creatorToken := models.Token{
		UserId:          uint(userId),
		ContractVersion: "v@next",
		Name:            "NewCoin",
		Symbol:          "NU",
		Network:         "MAINNET",
		OwnerAddress:    "0x810aafa05789dd461f81b22de31a2b21b2fbe3be",
		Address:         address.String(),
		TxHash:          tx.Hash().String(),
	}
	err = db.Create(&creatorToken).Error
	if err != nil {
		return models.Token{}, err
	}
	return creatorToken, err
}


