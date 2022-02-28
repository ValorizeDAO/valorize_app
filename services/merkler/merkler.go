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
	sliceToHash := [][]byte{n.LeftBytes, n.RightBytes}
	sort.SliceStable(sliceToHash, func(i, j int) bool {
		return bytes.Compare(sliceToHash[i], sliceToHash[j]) <= 0
	})
	hash := solsha3.SoliditySHA3(
		solsha3.Bytes32(sliceToHash[0]),
		solsha3.Bytes32(sliceToHash[1]),
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
	nodes := turnLeavesIntoNodes(leaves)
	tree := makeTree(nodes)
	for i := 0; i < len(tree); i++ {
		fmt.Println("\n==\nTree: \n", tree[i], "\n==")
	}
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

func makeTree(base []Node) []Node {
	var nodes []Node
	for i := 0; i < len(base); i = i + 2 {
		if i + 1 < len(base) {
			node := Node{Left: base[i], LeftBytes: base[i].Hash(), Right: base[i+1], RightBytes: base[i+1].Hash()}
			nodes = append(nodes, node)
			fmt.Printf(
				"%v\n==\nNode: \nleft: %v\nright %v\nhash: %v\n==\n",
				i,
				hex.EncodeToString(base[i].Hash()),
				hex.EncodeToString(base[i+1].Hash()),
				hex.EncodeToString(node.Hash()),
			)
		} else {
			node := Node{Left: base[i], LeftBytes: base[i].Hash(), Right:  EmptyNode{}}
			nodes = append(nodes, node)
			fmt.Printf(
				"\n==\nNode: \nleft: %v\nright: null\nhash: %v\n==\n",
				hex.EncodeToString(base[i].Hash()),
				hex.EncodeToString(node.Hash()),
			)
		}
		sort.SliceStable(nodes, func(i, j int) bool {
			return bytes.Compare(nodes[i].Hash(), nodes[j].Hash()) > 0
		})
	}
	if len(nodes) == 1 {
		return nodes
	} else {
		return makeTree(nodes)
	}
}