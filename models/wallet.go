package models

import "github.com/jinzhu/gorm"

type Wallet struct {
	ID         uint   `gorm:"primary_key"`
	Address    string `json:"address" gorm:"type:varchar(1000);"`
	Raw        string  `gorm:"size:1000"`
	UserId     uint
	User	   User    `gorm:"foreignKey:UserId"`
}

func AddExternalWalletForUser(user *User, address string, db gorm.DB) error {
	wallet := Wallet {
		Address: address,
		User: *user,
	}
	return db.Save(&wallet).Error
}
