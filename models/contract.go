package models

import "github.com/jinzhu/gorm"

type Contract struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:text"`
	Ticker   string `json:"ticker" gorm:"type:text"`
	ABI      string `json:"ABI" gorm:"type:json"`
	UserID   uint
	User     User   `gorm:"foreignkey:UserID"`
	Deployed string `json:"deployed" gorm:"type:bool"`
	Address  string `json:"address" gorm:"type:text"`
}
