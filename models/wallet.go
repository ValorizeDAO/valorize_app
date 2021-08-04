package models

type Wallet struct {
	ID         uint   `gorm:"primary_key"`
	Address    string `json:"address" gorm:"type:varchar(1000);"`
	Raw        string  `gorm:"size:1000"`
	UserId     uint
	User			 User    `gorm:"foreignKey:UserId"`
}
