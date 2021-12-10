package main

import (
	"strconv"
	"time"
	"valorize-app/config"
	"valorize-app/db"
	"valorize-app/models"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func main() {
	cfg := config.NewConfig()
	database := db.Init(cfg)
	m := GetMigrations(database)
	err := m.Migrate()
	if err == nil {
		print("Migrations did run successfully, seeding now")
	} else {
		print("migrations failed.", err)
	}
}
func GetMigrations(database *gorm.DB) *gormigrate.Gormigrate {
	timeNow := strconv.FormatUint(uint64(time.Now().Unix()), 10)
	return gormigrate.New(database, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: timeNow + "_migration",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.AutoMigrate(&models.User{}).Error; err != nil {
					return err
				}
				if err := tx.AutoMigrate(&models.Wallet{}).
					AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").
					Error; err != nil {
					return err
				}
				if err := tx.AutoMigrate(&models.Token{}).
					AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").
					Error; err != nil {
					return err
				}
				if err := tx.AutoMigrate(&models.Link{}).
					AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").
					Error; err != nil {
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.DropTable("users").Error; err != nil {
					return nil
				}
				if err := tx.DropTable("wallets").Error; err != nil {
					return nil
				}
				if err := tx.DropTable("tokens").Error; err != nil {
					return nil
				}
				if err := tx.DropTable("links").Error; err != nil {
					return nil
				}
				return nil
			},
		},
	})
}
