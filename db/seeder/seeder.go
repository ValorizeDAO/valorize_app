package main

import (
	"fmt"
	"valorize-app/config"
	"valorize-app/db"
	"valorize-app/models"
	"valorize-app/services/ethereum"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	database := db.Init(config.NewConfig())
	err := RunSeeder(database)
	if err == nil {
		print("seeding did run successfully")
	} else {
		print("seeding failed.", err)
	}
}

func RunSeeder(db *gorm.DB) error {
	user, _ := UserSeeder(db, "javier123454321", "test@test.com", "test", true)
	UserSeeder(db, "zselyke", "test@test2.com", "test", false)
	creatorToken := models.Token{
		UserId:          user.ID,
		ContractVersion: "v0.0.1",
		Name:            "JAVICoin",
		Symbol:          "JABA",
		OwnerAddress:    "0x810aafa05789dd461f81b22de31a2b21b2fbe3be",
		Address:         "0x05F3074138b9bfAbe8ADf0f68f2AA33047a8192e",
		TxHash:          "0xb57e816e5ca3c68d7d9d7522380318e8fde8dd320a247ac15997c3cb0956553f",
	}
	err := db.Create(&creatorToken).Error
	return err
}

func UserSeeder(db *gorm.DB, username string, email string, password string, hasDeployed bool) (models.User, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	user := models.User{
		Email:            email,
		Username:         username,
		Password:         string(hash),
		HasDeployedToken: hasDeployed,
	}
	err := db.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	_, err = ethereum.StoreUserKeystore(password, user.ID, db)
	if err != nil {
		fmt.Println("error: " + err.Error())
		return models.User{}, err
	}
	return user, err
}
