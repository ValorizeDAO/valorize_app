package models

type Wallet struct {
	ID         uint   `gorm:"primary_key"`
	Address    string `json:"address" gorm:"type:varchar(1000);"`
	IsContract bool
	UserId     uint
}
