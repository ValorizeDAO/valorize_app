package models

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

type getAllAirdropClaimsTestIsRightLength struct {
	airdropId, expectedLength int
}

func TestGetAllAirdropClaimsIsRightLength(t *testing.T) {
	prepareTestDatabase()
	tests := []getAllAirdropClaimsTestIsRightLength{
		{airdropId: 1, expectedLength: 3},
		{airdropId: 2, expectedLength: 2},
	}
	for _, test := range tests {
		output, err := GetAllAirdropClaims(*gormDb, test.airdropId)
		if err != nil {
			t.Error(err)
		}
		if len(output) != test.expectedLength {
			t.Errorf("\nGetAllAirdropClaims() returns all: returned %v, want %v", len(output), test.expectedLength)
		}
	}
}

type getAllAirdropClaimsTest struct {
	airdropId int
	expected  []AirdropClaim
}

func TestGetAllAirdropClaimsReturnsCorrectValues(t *testing.T) {
	prepareTestDatabase()
	tests := []getAllAirdropClaimsTest{
		{1, []AirdropClaim{
			{WalletAddress: "0x0", ClaimAmount: "100000", AirdropID: 1},
			{WalletAddress: "0x1", ClaimAmount: "200000", AirdropID: 1},
			{WalletAddress: "0x2", ClaimAmount: "200000", AirdropID: 1},
		}},
		{2, []AirdropClaim{
			{WalletAddress: "0x1", ClaimAmount: "200000", AirdropID: 2},
			{WalletAddress: "0x2", ClaimAmount: "400000", AirdropID: 2},
		}},
	}
	for _, test := range tests {
		allClaims, err := GetAllAirdropClaims(*gormDb, test.airdropId)
		if err != nil {
			t.Error(err)
		}
		for i, claim := range allClaims {
			if test.expected[i] != claim {
				t.Errorf("\nGetAllAirdropClaims() returns correct values for Airdrop ID %v: returned %v, want %v", test.airdropId, claim, test.expected)
			}
		}
	}
}

type getLatestAirdropTest struct {
	tokenId      int
	onChainIndex int
	expected     Airdrop
}

func TestGetLatestAirdropByTokenId(t *testing.T) {
	prepareTestDatabase()
	tests := []getLatestAirdropTest{
		{2, 1, Airdrop{ID: 1, TokenID: 2, OnChainIndex: 1, MerkleRoot: "0xRoot"}},
		{3, 2, Airdrop{ID: 4, TokenID: 3, OnChainIndex: 2, MerkleRoot: "0xRoot34"}},
		{5, 10, Airdrop{ID: 3, TokenID: 5, OnChainIndex: 10, MerkleRoot: "0xRoot4"}},
	}
	for _, test := range tests {
		airdrop, err := GetAirdropByTokenIndexAndOnChainId(*gormDb, test.tokenId, test.onChainIndex)
		if err != nil {
			t.Error(err)
		}
		if airdrop != test.expected {
			t.Errorf("\nGetAirdropByTokenIndexAndOnChainId() returns correct values for token ID %v, onChainIndex %v: returned %v, want %v", test.tokenId, test.onChainIndex, airdrop, test.expected)
		}
	}
}
