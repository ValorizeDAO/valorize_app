package main

import (
	"flag"
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

	drop := flag.String("d", "none", "drop token_type column")
	flag.Parse()
	m := GetMigrations(database)

	if *drop == "token_type" {
		database.Exec("ALTER TABLE tokens DROP COLUMN token_type;")
	}

	err := m.Migrate()

	if err == nil {
		print("Migrations did run successfully")
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
				if err := tx.AutoMigrate(&models.AirdropClaim{}).
					AddForeignKey("token_id", "tokens(id)", "CASCADE", "CASCADE").
					Error; err != nil {
					return err
				}
				database.Exec("ALTER TABLE airdrop_claims MODIFY COLUMN claimed tinyint(1) DEFAULT 0 NULL;")
				database.Exec("ALTER TABLE airdrop_claims CHANGE claim_amount claim_amount varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NUL;")
				database.Exec("ALTER TABLE airdrop_claims CHANGE address wallet_address  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL;")

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
				if err := tx.DropTable("airdrop_claims").Error; err != nil {
					return nil
				}
				return nil
			},
		},
	})
}
