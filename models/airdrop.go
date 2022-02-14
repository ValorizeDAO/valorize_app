package models

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

type Airdrop struct {
	ID            uint   `json:"id" gorm:"primary_key"`
	TokenID       uint   `json:"token_id"`
	MerkleRoot    string `json:"merkle_root"`
	RawMerkleTree string `json:"raw_merkle_tree" gorm:"type:longtext;"`
	FinishTime    uint   `json:"finish_time" gorm:"not null"`
	Complete      bool   `json:"complete" gorm:"not null" sql:"DEFAULT:0"`
}

type AirdropClaim struct {
	WalletAddress string  `json:"wallet_address"`
	ClaimAmount   string  `json:"claim_amount"`
	AirdropID     uint    `json:"airdrop_id"`
	Airdrop       Airdrop `json:"airdrop"`
	Claimed       bool    `json:"claimed" gorm:"not null" sql:"DEFAULT:0"`
}

func NewAirdropClaim(db gorm.DB, claimInfo [][]string, tokenId int, airdropId uint) error {
	chunkSize := 150
	if len(claimInfo) < chunkSize {
		chunkSize = len(claimInfo)
	}
	tx := db.Begin()
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
