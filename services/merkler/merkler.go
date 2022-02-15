package merkler

import (
	"log"

	"github.com/cbergoon/merkletree"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

type ContentData struct {
	Address     string
	ClaimAmount string
}

//CalculateHash hashes the values of a TestContent
func (c ContentData) CalculateHash() ([]byte, error) {
	hash := solsha3.SoliditySHA3(
		[]string{"address", "uint256"},
		[]interface{}{
			c.Address,
			c.ClaimAmount,
		},
	)
	return hash, nil
}
func (c ContentData) Equals(other merkletree.Content) (bool, error) {
	hashThis := solsha3.SoliditySHA3(
		[]string{"address", "uint256"},
		[]interface{}{
			c.Address,
			c.ClaimAmount,
		},
	)
	hashOther := solsha3.SoliditySHA3(
		[]string{"address", "uint256"},
		[]interface{}{
			c.Address,
			c.ClaimAmount,
		},
	)
	return string(hashThis) == string(hashOther), nil
}

func GenerateAirdropMerkleRoot(addressToClaimAmountMap [][]string) []byte {
	var combinedStrings []merkletree.Content
	for i := 0; i < len(addressToClaimAmountMap)-1; i++ {
		combinedStrings = append(
			combinedStrings,
			ContentData{addressToClaimAmountMap[i][0], addressToClaimAmountMap[i][1]},
		)
	}
	t, err := merkletree.NewTree(combinedStrings)
	if err != nil {
		log.Fatal(err)
	}
	mr := t.MerkleRoot()
	log.Println(mr)
	return mr
}
