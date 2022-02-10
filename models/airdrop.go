package models

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

type AirdropClaim struct {
	TokenId       uint   `json:"token_id"`
	WalletAddress string `json:"wallet_address"`
	ClaimAmount   string `json:"claim_amount"`
	AirdropId     uint   `json:"airdrop_id"`
	Claimed       bool   `json:"claimed"; gorm:"not null"; default:0;`
}

func NewAirdropClaim(db gorm.DB, claimInfo [][]string, tokenId int, airdropId uint) error {
	var chunkSize int
	if len(claimInfo) >= 100 {
		chunkSize = 100
	} else {
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
			valueStrings = append(valueStrings, "(?, ?, ?, ?)")
			valueArgs = append(valueArgs, tokenId)
			valueArgs = append(valueArgs, claim[0])
			valueArgs = append(valueArgs, claim[1])
			valueArgs = append(valueArgs, airdropId)
		}

		stmt := fmt.Sprintf("INSERT INTO airdrop_claims (token_id, wallet_address, claim_amount, airdrop_id) VALUES %s", strings.Join(valueStrings, ","))
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
