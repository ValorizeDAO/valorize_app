package models

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

type Airdrop struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	TokenID      uint   `json:"token_id"`
	MerkleRoot   string `json:"merkle_root"`
	OnChainIndex uint   `json:"onchain_index"` //The contract has an index which is stored in this field
}

type AirdropClaim struct {
	WalletAddress string  `json:"wallet_address"`
	ClaimAmount   string  `json:"claim_amount"`
	AirdropID     uint    `json:"airdrop_id"`
	Airdrop       Airdrop `json:"airdrop"`
	Claimed       bool    `json:"claimed" gorm:"not null" sql:"DEFAULT:0"`
}

func (m *Model) GetAirdropByTokenId(tokenId uint64) (Airdrop, error) {
	var a Airdrop
	if err := m.db.Where("token_id=?", tokenId).Order("on_chain_index desc").First(&a).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return Airdrop{}, err
		}
		return Airdrop{}, err
	}
	return a, nil
}

func (m *Model) GetAirdropClaimByAddress(walletaddress string) (AirdropClaim, error) {
	var ac AirdropClaim
	if err := m.db.Preload("Airdrop").Where("wallet_address=?", walletaddress).First(&ac).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return AirdropClaim{}, err
		}
		return AirdropClaim{}, err
	}
	return ac, nil
}

func (m *Model) GetAirdropClaimByAirdropID(airdropId uint64, walletaddress string) (AirdropClaim, error) {
	var ac AirdropClaim
	if err := m.db.Where("airdrop_id=?", airdropId).Where("wallet_address=?", walletaddress).Find(&ac).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return AirdropClaim{}, err
		}
		return AirdropClaim{}, err
	}
	return ac, nil
}

func (m *Model) NewAirdrop(drop Airdrop) (Airdrop, error) {
	if err := m.db.Save(&drop).Error; err != nil {
		return Airdrop{}, err
	}
	return drop, nil
}

func (m *Model) NewAirdropClaim(claimInfo [][]string, airdropId uint) error {
	chunkSize := 150
	if len(claimInfo) < chunkSize {
		chunkSize = len(claimInfo)
	}
	tx := m.db.Begin()
	var divided [][][]string

	// grabs the first { size } and makes a list of chunks
	for i := 0; i < len(claimInfo)-1; i += chunkSize {
		end := i + chunkSize

		if end > len(claimInfo) {
			end = len(claimInfo)
		}
		divided = append(divided, claimInfo[i:end])

		//create an array to be converted to insert string
		valueStrings := []string{}
		valueArgs := []interface{}{}
		for _, claim := range divided[len(divided)-1] {
			valueStrings = append(valueStrings, "(?, ?, ?)")
			valueArgs = append(valueArgs, claim[0])
			valueArgs = append(valueArgs, claim[1])
			valueArgs = append(valueArgs, airdropId)
		}

		stmt := fmt.Sprintf("INSERT INTO airdrop_claims (wallet_address, claim_amount, airdrop_id) VALUES %s", strings.Join(valueStrings, ","))
		err := tx.Exec(stmt, valueArgs...).Error
		if err != nil {
			tx.Rollback()
			fmt.Println(err)
			return err
		}
	}
	err := tx.Commit().Error
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func (m *Model) GetAirdropByTokenIndexAndOnChainId(tokenId int, onChainIndex int) (Airdrop, error) {
	var airdrop Airdrop
	err := m.db.Where("token_id = ? AND on_chain_index = ?", tokenId, onChainIndex).Order("id desc").First(&airdrop).Error
	if err != nil {
		return Airdrop{}, err
	}
	return airdrop, nil
}

func (m *Model) GetAllAirdropClaims(airdrop_id int) ([]AirdropClaim, error) {
	var claims []AirdropClaim
	err := m.db.Where("airdrop_id = ?", airdrop_id).Find(&claims).Error
	if err != nil {
		return nil, err
	}
	return claims, nil
}
