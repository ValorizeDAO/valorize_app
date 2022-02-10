package models

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

type AirdropClaim struct {
	TokenId     uint   `json:"token_id"`
	Address     string `json:"wallet_address"`
	ClaimAmount string `json:"claim_amount"`
	AirdropId   uint   `json:"airdrop_id"`
	Claimed     bool   `json:"claimed"; gorm:"not null"; default:0;`
}

func NewAirdropClaim(db gorm.DB, claimInfo [][]string, tokenId int, airdropId uint) error {
	size := 500
	tx := db.Begin()
	var divided [][][]string

	chunkSize := (len(claimInfo) + size) / size
	// grabs the first { size } and makes a list of chunks
	for i := 0; i < len(claimInfo); i += chunkSize {
		end := i + chunkSize

		if end > len(claimInfo) {
			end = len(claimInfo)
		}
		divided = append(divided, claimInfo[i:end])

		valueStrings := []string{}
		valueArgs := []interface{}{}
		for _, claim := range divided[i] {
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
	return nil
}
