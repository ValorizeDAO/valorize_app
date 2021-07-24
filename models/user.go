package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"type:varchar(200);"`
	Name     string `json:"name" gorm:"type:varchar(200);"`
	Username string `json:"username" gorm:"type:varchar(200);"`
	Password string `json:"password" gorm:"type:varchar(200);"`
	Wallets  []Wallet
}

func GetUserByID(id uint, db gorm.DB) (*User, error) {
	var m User
	if err := db.First(&m, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func GetUserByEmail(e string, db gorm.DB) (*User, error) {
	var m User
	if err := db.Where(&User{Email: e}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func GetUserByUsername(username string, db gorm.DB) (*User, error) {
	var m User
	if err := db.Where(&User{Username: username}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}
