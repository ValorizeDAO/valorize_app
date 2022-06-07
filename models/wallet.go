package models

import (
	"errors"
)

type Wallet struct {
	ID             uint     `gorm:"primary_key"`
	Address        string   `json:"address" gorm:"type:varchar(1000);"`
	Raw            string   `gorm:"size:1000"`
	UserId         uint     `gorm:"default:null;"`
	User           User     `gorm:"foreignKey:UserId"`
	Administrating []*Token `json:"is_administrating" gorm:"type:varchar(200); many2many:token_wallet;"`
}

func (m *Model) AddExternalWalletForUser(user *User, address string) error {
	wallet := Wallet{}
	if m.db.Where("address = ? AND user_id = ?", address, user.ID).Find(&wallet).Error != nil {
		wallet := Wallet{
			Address: address,
			User:    *user,
		}
		return m.db.Save(&wallet).Error
	}
	return errors.New(user.Username + " already has a wallet with address " + address)
}

func (m *Model) GetAllWalletsByUserId(userId uint) ([]string, error) {
	wallets := []Wallet{}
	err := m.db.Select("address").Where("user_id = ?", userId).Find(&wallets).Error
	addresses := make([]string, len(wallets))
	for i, wallet := range wallets {
		addresses[i] = wallet.Address
	}
	return addresses, err
}

func (m *Model) GetWalletDataFromAddress(address string) Wallet {
	wallet := Wallet{}
	m.db.FirstOrCreate(&wallet, Wallet{Address: address})
	return wallet
}
