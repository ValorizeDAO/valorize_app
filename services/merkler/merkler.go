package merkler

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"valorize-app/services/ethereum"

	"github.com/cbergoon/merkletree"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

type ContentData struct {
	data []byte
}

func (c ContentData) CalculateHash() ([]byte, error) {
	return solsha3.SoliditySHA3(solsha3.Bytes1(c.data)), nil
}

func (c ContentData) Equals(other merkletree.Content) (bool, error) {
	otherHash, _ := other.CalculateHash()
	thisHash, _ := c.CalculateHash()
	return hex.EncodeToString(thisHash) == hex.EncodeToString(otherHash), nil
}

func GenerateAirdropMerkleRoot(addressToClaimAmountMap [][]string) ([]byte, error) {

	var leaves [][]byte
	for i := 0; i < len(addressToClaimAmountMap); i++ {
		address := addressToClaimAmountMap[i][0]
		claimString := addressToClaimAmountMap[i][1]
		claimAmount, ok := new(big.Int).SetString(claimString, 10)
		if !ok || !ethereum.CheckAddress(address) {
			return []byte(""), errors.New("invalid value in input for address: " + address + ", claimAmount: " + claimString + "in index: " + strconv.Itoa(i))
		}

		hash := solsha3.SoliditySHA3(
			solsha3.Address(address),
			solsha3.Uint256(claimAmount),
		)

		leaves = append(leaves, hash)
		
		fmt.Printf("0x%v\n", hex.EncodeToString(hash))
	}
	// sort.Slice(leaves, func(i, j int) bool {
	// 	return bytes.Compare(leaves[i], leaves[j]) < 0
	// })
	var contents []merkletree.Content
	for i := 0; i < len(leaves); i++ {
		contents = append(contents, ContentData{leaves[i]})
	}
	t, err := merkletree.NewTree(contents)
	if err != nil {
		return []byte(""), err
	}
	mr := t.MerkleRoot()
	return mr, nil
}