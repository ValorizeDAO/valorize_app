package merkler

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"valorize-app/services/ethereum"

	"github.com/cbergoon/merkletree"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type ContentData struct {
	ByteArray []byte
}

func (c ContentData) CalculateHash() ([]byte, error) {
	return crypto.Keccak256(c.ByteArray), nil
}

func (c ContentData) Equals(other merkletree.Content) (bool, error) {
	otherHash, _ := other.CalculateHash()
	thisHash, _ := c.CalculateHash()
	return hex.EncodeToString(thisHash) == hex.EncodeToString(otherHash), nil
}

func GenerateAirdropMerkleRoot(addressToClaimAmountMap [][]string) ([]byte, error) {
	var leaves []merkletree.Content
	for i := 0; i < len(addressToClaimAmountMap)-1; i++ {
		address := addressToClaimAmountMap[i][0]
		claimString := addressToClaimAmountMap[i][1]
		claimAmount, ok := new(big.Int).SetString(claimString, 10)
		if !ok || !ethereum.CheckAddress(address) {
			return []byte(""), errors.New("invalid value in input for address: " + address + ", claimAmount: " + claimString + "in index: " + strconv.Itoa(i))
		}
		addressType, _ := abi.NewType("address", "", nil)
		uintType, _ := abi.NewType("uint256", "", nil)
		arguments := abi.Arguments{
			{
				Type: addressType,
			},
			{
				Type: uintType,
			},
		}

		bytes, _ := arguments.Pack(
			common.HexToAddress(address),
			claimAmount,
		)

		leaves = append(
			leaves, ContentData{bytes},
		)
		fmt.Printf("\n0x%v", hex.EncodeToString(crypto.Keccak256(bytes)))
	}
	t, err := merkletree.NewTree(leaves)
	if err != nil {
		return []byte(""), err
	}
	mr := t.MerkleRoot()
	return mr, nil
}
