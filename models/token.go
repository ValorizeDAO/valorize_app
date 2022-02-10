package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Token struct {
	ID              uint       `json:"id" gorm:"primary_key"`
	TokenType       string     `json:"token_type" gorm:"not null; type:enum('simple', 'creator', 'timed_mint'); default:'creator';"`
	Name            string     `json:"name" gorm:"type:varchar(200);"`
	Address         string     `json:"address" gorm:"type:varchar(200);"`
	OwnerAddress    string     `json:"owner_address" gorm:"type:varchar(200);"`
	VaultAddress    string     `json:"vault_address" gorm:"type:varchar(200);"`
	AdminAddresses  []*Wallet  `json:"administrated_by" gorm:"type:varchar(200); many2many:token_wallet;"`
	Symbol          string     `json:"symbol" gorm:"type:varchar(200);"`
	ChainId         string     `json:"chain_id" gorm:"type:varchar(200);"`
	ContractVersion string     `json:"contract_version" gorm:"type:varchar(200);"`
	TxHash          string     `json:"tx_hash" gorm:"type:varchar(200);"`
	UserId          uint       `json:"user_id"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at" sql:"index"`
}

type TokenResponse struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	Address         string `json:"address"`
	OwnerAddress    string `json:"owner_address"`
	VaultAddress    string `json:"vault_address"`
	Symbol          string `json:"symbol"`
	ContractVersion string `json:"contract_version"`
	TxHash          string `json:"tx_hash"`
	UserId          uint   `json:"user_id"`
	ChainId         string `json:"chain_id"`
	TokenType       string `json:"token_type"`
}

func GetTokenResponse(token *Token) TokenResponse {
	return TokenResponse{
		ID:              token.ID,
		Name:            token.Name,
		Address:         token.Address,
		OwnerAddress:    token.OwnerAddress,
		VaultAddress:    token.VaultAddress,
		Symbol:          token.Symbol,
		ContractVersion: token.ContractVersion,
		TxHash:          token.TxHash,
		UserId:          token.UserId,
		ChainId:         token.ChainId,
		TokenType:       token.TokenType,
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
