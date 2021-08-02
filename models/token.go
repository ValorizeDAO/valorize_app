package models

import "github.com/jinzhu/gorm"

type Token struct {
  gorm.Model
  Name            string `json:"name" gorm:"type:varchar(200);"`
  OwnerAddress    string `json:"username" gorm:"type:varchar(200);"`
  Symbol          string `json:"password" gorm:"type:varchar(200);"`
  Network         string `json:"Avatar" gorm:"type:varchar(200);"`
  About           string `json:"About" gorm:"type:varchar(1000);"`
  UserId          uint
  User            User  `gorm:"foreignKey:UserId"`
}
