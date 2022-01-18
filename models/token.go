package models

import (
	"database/sql/driver"
	"time"

	"github.com/jinzhu/gorm"
)

type tokenType string

const (
	creator tokenType = "creator"
	simple  tokenType = "simple"
)

func (t *tokenType) Scan(value interface{}) error {
	*t = tokenType(value.([]byte))
	return nil
}

func (t tokenType) Value() (driver.Value, error) {
	return string(t), nil
}

type Token struct {
	ID              uint       `json:"id" gorm:"primary_key"`
	TokenType       string     `json:"token_type2" gorm:"not null; type:enum('simple', 'creator'); default:'creator';"`
	Name            string     `json:"name" gorm:"type:varchar(200);"`
	Address         string     `json:"address" gorm:"type:varchar(200);"`
	OwnerAddress    string     `json:"owner_address" gorm:"type:varchar(200);"`
	VaultAddress    string     `json:"vault_address" gorm:"type:varchar(200);"`
	AdminAddresses  []*Wallet  `json:"token_admin" gorm:"type:varchar(200); many2many:token_admin;"`
	Symbol          string     `json:"symbol" gorm:"type:varchar(200);"`
	Network         string     `json:"network" gorm:"type:varchar(200);"`
	ChainId         string     `json:"chain_id" gorm:"type:varchar(200);"`
	ContractVersion string     `json:"contract_version" gorm:"type:varchar(200);"`
	TxHash          string     `json:"tx_hash" gorm:"type:varchar(200);"`
	UserId          uint       `json:"user_id"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at" sql:"index"`
}

func (Token) TableName() string {
	return "tokens"
}

type TokenResponse struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	Address         string `json:"address"`
	OwnerAddress    string `json:"owner_address"`
	Symbol          string `json:"symbol"`
	Network         string `json:"network"`
	ContractVersion string `json:"contract_version"`
	UserId          uint   `json:"user_id"`
	TxHash          string `json:"tx_hash"`
}

func GetTokenResponse(creatorToken *Token) TokenResponse {
	return TokenResponse{
		ID:              creatorToken.ID,
		Name:            creatorToken.Name,
		Address:         creatorToken.Address,
		OwnerAddress:    creatorToken.OwnerAddress,
		Symbol:          creatorToken.Symbol,
		Network:         creatorToken.Network,
		ContractVersion: creatorToken.ContractVersion,
		UserId:          creatorToken.UserId,
	}
}

func GetTokenById(tokenId uint64, db gorm.DB) (TokenResponse, error) {
	var t Token
	if err := db.Where("id=?", tokenId).First(&t).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return TokenResponse{}, err
		}
		return TokenResponse{}, err
	}
	return GetTokenResponse(&t), nil
}
