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
	// send username parameter to command line
	username := os.Args[1]
	fmt.Print(os.Args)
	user, err := AddToAlpha(database, username)
	if err == nil {
		fmt.Print("Added to alpha ", user.Username)
	}

}

func AddToAlpha(db *gorm.DB, username string) (models.User, error) {
	m := models.NewModels(db)
	userData, err := m.GetUserByUsername(username)

	if err != nil {
		fmt.Println("could not find user " + username + err.Error())
		return models.User{}, err
	}
	userData.IsAlphaUser = true

	err = db.Save(&userData).Error
	if err != nil {
		fmt.Println(err.Error())
		return models.User{}, err
	}
	return userData, nil
}
