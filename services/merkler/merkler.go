package merkler

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"valorize-app/services/ethereum"

	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

type ContentData struct {
	data []byte
}

// func (c ContentData) CalculateHash() ([]byte, error) {
// 	return crypto.Keccak256(solsha3.Bytes32(c.data)), nil
// }

// func (c ContentData) Equals(other merkletree.Content) (bool, error) {
// 	otherHash, _ := other.CalculateHash()
// 	thisHash, _ := c.CalculateHash()
// 	return hex.EncodeToString(thisHash) == hex.EncodeToString(otherHash), nil
// }


type Node struct {
	Left 		Hashable
	Right 		Hashable
	LeftBytes 	[]byte
	RightBytes 	[]byte
}
type Leaf struct {
	Address string
	ClaimAmount *big.Int
}
type BaseHashable struct {
	BaseBytes []byte
}

func (base BaseHashable) Hash() []byte {
	fmt.Println("Hashing base bytes")
	return solsha3.SoliditySHA3(
	   solsha3.Bytes32(base.BaseBytes),
   )
}

func (leaf Leaf) Hash() []byte {
	hash := solsha3.SoliditySHA3(
		solsha3.Address(leaf.Address),
		solsha3.Uint256(leaf.ClaimAmount),
	)
	return hash
}

func (leaf Leaf) ToHashableBytes() BaseHashable {
	return BaseHashable{BaseBytes: leaf.Hash()}
}

type Hash [20]byte

type Hashable interface {
	Hash() []byte
}

func (n Node) Hash() []byte {
	hash := solsha3.SoliditySHA3(
		solsha3.Bytes32(n.LeftBytes),
		solsha3.Bytes32(n.RightBytes),
	)
	fmt.Printf("\n==\nnode(inside Hash()): \nleft: %v\nright %v\n hash: %v\n==\n", 
		hex.EncodeToString(n.LeftBytes), 
		hex.EncodeToString(n.RightBytes), 
		hex.EncodeToString(hash),
	)
 	return hash
}
type EmptyNode []byte

func (e EmptyNode) Hash() []byte {
	return []byte{}
}

func GenerateAirdropMerkleRoot(addressToClaimAmountMap [][]string) ([]byte, error) {
	var leaves []Leaf
	for i := 0; i < len(addressToClaimAmountMap); i++ {
		address := addressToClaimAmountMap[i][0]
		claimString := addressToClaimAmountMap[i][1]
		claimAmount, ok := new(big.Int).SetString(claimString, 10)
		if !ok || !ethereum.CheckAddress(address) {
			return []byte(""), errors.New("invalid value in input for address: " + address + ", claimAmount: " + claimString + "in index: " + strconv.Itoa(i))
		}
		leaves = append(leaves, Leaf{Address: address, ClaimAmount: claimAmount})
	}
	sort.SliceStable(leaves, func(i, j int) bool {
		return bytes.Compare(leaves[i].Hash(), leaves[j].Hash()) < 0
	})
	_ = NewMerkleTree(leaves)
	return []byte{}, nil
}

func NewMerkleTree(leaves []Leaf) error {
	_ = turnLeavesIntoNodes(leaves)
	return nil
}

func turnLeavesIntoNodes(leaves []Leaf) []Node {
	var nodes []Node
	for i := 0; i < len(leaves); i = i + 2 {
		if i + 1 < len(leaves) {
			nodes = append(nodes, Node{
				Left: leaves[i].ToHashableBytes(), 
				Right: leaves[i+1].ToHashableBytes(),
				LeftBytes: leaves[i].Hash(),
				RightBytes: leaves[i+1].Hash(),
			})
		} else {
			nodes = append(nodes, Node{Left: leaves[i].ToHashableBytes(), Right: EmptyNode{}, LeftBytes: leaves[i].Hash()})
		}
	}
	return nodes
}


// func makeTree(base []Hashable) []Hashable {
// 	var nodes []Hashable
// 	for i := 0; i < len(base); i = i + 2 {
// 		if i + 1 < len(base) {
// 			nodes = append(nodes, Node{Left: base[i], Right: base[i+1]})
// 		} else {
// 			nodes = append(nodes, Node{Left: base[i], Right: EmptyNode{}})
// 		}
// 		sort.SliceStable(nodes, func(i, j int) bool {
// 			return bytes.Compare(nodes[i].Hash(), nodes[j].Hash()) < 0
// 		})
// 	}
// 	if len(nodes) == 1 {
// 		return nodes
// 	} else {
// 		return makeTree(nodes)
// 	}
// }