package models

import (
  "time"
)

type Token struct {
  ID              uint       `json:"id" gorm:"primary_key"`
  Name            string     `json:"name" gorm:"type:varchar(200);"`
  Address         string     `json:"address" gorm:"type:varchar(200);"`
  OwnerAddress    string     `json:"owner_address" gorm:"type:varchar(200);"`
  Symbol          string     `json:"symbol" gorm:"type:varchar(200);"`
  Network         string     `json:"network" gorm:"type:varchar(200);"`
  ContractVersion string     `json:"contract_version" gorm:"type:varchar(200);"`
  TxHash          string     `json:"tx_hash" gorm:"type:varchar(200);"`
  UserId          uint       `json:"user_id"`
  User            User       `json:"user" gorm:"foreignKey:UserId;AssociationForeignKey:ID"`
  CreatedAt       time.Time  `json:"created_at"`
  UpdatedAt       time.Time  `json:"updated_at"`
  DeletedAt       *time.Time `json:"deleted_at" sql:"index"`
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
