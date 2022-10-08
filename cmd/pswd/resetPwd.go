package main

import (
	"fmt"
	"os"
	"valorize-app/config"
	"valorize-app/db"
	"valorize-app/models"
	"valorize-app/services"

	"github.com/jinzhu/gorm"
)

func main() {
	cfg := config.NewConfig()
	database := db.Init(cfg)
	defer database.Close()
	// send email parameter to command line
	email := os.Args[1]
	fmt.Print(os.Args)
	link, duration, err := GeneratePwdResetLink(database, email)
	if err == nil {
		fmt.Printf("Generated Password Reset Link:\n%v\nactive until: %v", link, duration)
	}

}

func GeneratePwdResetLink(db *gorm.DB, email string) (string, int64, error) {
	m := models.NewModels(db)
	user, err := m.GetUserByEmail(email)

	if err != nil {
		fmt.Println("could not find user " + email + err.Error())
		return "", 0, err
	}

	jwt, duration, err := services.NewToken(user, 1)
	if err != nil {
		fmt.Println(err.Error())
		return "", 0, err
	}

	return os.Getenv("FRONTEND_URL") + "/reset-password?credential=" + jwt, duration, nil
}
