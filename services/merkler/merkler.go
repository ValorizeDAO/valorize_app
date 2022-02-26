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
		
	}
	sort.SliceStable(leaves, func(i, j int) bool {
		return bytes.Compare(leaves[i], leaves[j]) < 0
	})
	// var contents []merkletree.Content
	// for i := 0; i < len(leaves); i++ {
	// 	contents = append(contents, ContentData{leaves[i]})
	// }
	// var hashes []string
	// for i := 0; i < len(contents); i++ {
	// 	content, _ := contents[i].CalculateHash()
	// 	hashes = append(hashes, "0x" + hex.EncodeToString(content) + ",\n")
	// }
	_ = NewMerkleTree(leaves, solsha3.Bytes32)
	// fmt.Println(hashes)
	// if err != nil {
	// 	return []byte(""), err
	// }
	// mr := t.MerkleRoot()
	return []byte{}, nil
}

func NewMerkleTree(leaves [][]byte, hashing func(interface{}) []byte) error {
	tree := makeTree(leaves)
	for _, leaf := range tree {
		fmt.Printf("leaf: \n%v\n", hex.EncodeToString(leaf))
	}
	return nil
}

func makeTree(leaves [][]byte) [][]byte {
	var branches [][]byte
	for i := 0; i < len(leaves); i = i + 2 {
		if (i + 1 == len(leaves)) {
			return append(branches, leaves[i])
		} else {
			left := leaves[i]
			right := leaves[i+1]
			hash := solsha3.SoliditySHA3(
				solsha3.Bytes32(left),
				solsha3.Bytes32(right),
			)
			fmt.Printf("Hashing: \n 0x%v w/ 0x%v\n result: \n0x%v\n\n",
			 hex.EncodeToString(left),
			 hex.EncodeToString(right),
			 hex.EncodeToString(hash),
			)
			branches = append(branches, hash)
			sort.SliceStable(branches, func(i, j int) bool {
				return bytes.Compare(branches[i], branches[j]) < 0
			})
		}
	}
	
	if len(branches[len(branches)-1]) != 1 {
		branches = makeTree(branches)
	}
	return branches
}
// _createHashes(nodes) {
// 	while (nodes.length > 1) {
// 		const layerIndex = this.layers.length;
// 		this.layers.push([]);
// 		for (let i = 0; i < nodes.length; i += 2) {
// 			if (i + 1 === nodes.length) {
// 				if (nodes.length % 2 === 1) {
// 					let data = nodes[nodes.length - 1];
// 					let hash = data;
// 					// is bitcoin tree
// 					if (this.isBitcoinTree) {
// 						// Bitcoin method of duplicating the odd ending nodes
// 						data = Buffer.concat([buffer_reverse_1.default(data), buffer_reverse_1.default(data)]);
// 						hash = this.hashFn(data);
// 						hash = buffer_reverse_1.default(this.hashFn(hash));
// 						this.layers[layerIndex].push(hash);
// 						continue;
// 					}
// 					else {
// 						if (this.duplicateOdd) {
// 							// continue with creating layer
// 						}
// 						else {
// 							// push copy of hash and continue iteration
// 							this.layers[layerIndex].push(nodes[i]);
// 							continue;
// 						}
// 					}
// 				}
// 			}
// 			const left = nodes[i];
// 			const right = i + 1 === nodes.length ? left : nodes[i + 1];
// 			let data = null;
// 			let combined = null;
// 			if (this.isBitcoinTree) {
// 				combined = [buffer_reverse_1.default(left), buffer_reverse_1.default(right)];
// 			}
// 			else {
// 				combined = [left, right];
// 			}
// 			if (this.sortPairs) {
// 				combined.sort(Buffer.compare);
// 			}
// 			data = Buffer.concat(combined);
// 			let hash = this.hashFn(data);
// 			// double hash if bitcoin tree
// 			if (this.isBitcoinTree) {
// 				hash = buffer_reverse_1.default(this.hashFn(hash));
// 			}
// 			this.layers[layerIndex].push(hash);
// 		}
// 		nodes = this.layers[layerIndex];
// 	}
// }