package models

import (
	"github.com/jinzhu/gorm"
)

type SmartContract struct {
	ID       uint   `gorm:"primary_key"`
	Key      string `json:"key" gorm:"unique;"`
	ByteCode string `json:"byte_code"`
}

func GetSmartContractByKey(key string, db gorm.DB) (SmartContract, error) {
	var contractInfo SmartContract
	err := db.Where("`key` = ?", key).First(&contractInfo).Error
	if err != nil {
		return SmartContract{}, err
	}
	return contractInfo, err
}
