package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Wallet struct {
	ID         uint   `gorm:"primary_key"`
	Address    string `json:"address" gorm:"type:varchar(1000);"`
	Raw        string  `gorm:"size:1000"`
	UserId     uint
	User	   User    `gorm:"foreignKey:UserId"`
}

func AddExternalWalletForUser(user *User, address string, db gorm.DB) error {
	wallet := Wallet{}
	if db.Where("address = ? AND user_id = ?", address, user.ID).Find(&wallet).Error != nil {
		wallet := Wallet {
			Address: address,
			User: *user,
		}
		return db.Save(&wallet).Error
	}
	return errors.New(user.Username + " already has a wallet with address " + address)
}


func GetAllWalletsByUserId(userId uint, db gorm.DB) ([]string, error) {
	wallets := []Wallet{}
	err := db.Select("address").Where("user_id = ?", userId).Find(&wallets).Error
	addresses := make([]string, len(wallets))
	for i, wallet := range wallets {
		addresses[i] = wallet.Address
	}
	return addresses, err
}