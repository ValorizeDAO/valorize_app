package main

import (
  "github.com/jinzhu/gorm"
  "gopkg.in/gormigrate.v1"
  "strconv"
  "time"
  "valorize-app/config"
  "valorize-app/db"
  "valorize-app/models"
)

func main() {
  cfg := config.NewConfig()
  m := GetMigrations(db.Init(cfg))
  err := m.Migrate()
  if err == nil {
    print("Migrations did run successfully")
  } else {
    print("migrations failed.", err)
  }
}
func GetMigrations(db *gorm.DB) *gormigrate.Gormigrate {
  timeNow := strconv.FormatUint(uint64(time.Now().Unix()), 10)
  return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
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
        return nil
      },
    },
  })
}
//func getMigrationsList() []gmStore.Migratable {
//	return []gmStore.Migratable{
//		&list.CreateUserTable{},
//		&list.CreateWalletTable{},
//		&list.UpdateUserTable{},
//		&list.CreateTokenTable{},
//	}
//}
