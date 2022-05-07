package models_test

import (
	"testing"
	"valorize-app/models"

	_ "github.com/go-sql-driver/mysql"
)

type getAllAirdropClaimsTestIsRightLength struct {
	airdropId, expectedLength int
}

func TestGetAllAirdropClaimsIsRightLength(t *testing.T) {
	prepareTestDatabase()
	testTable := []getAllAirdropClaimsTestIsRightLength{
		{airdropId: 1, expectedLength: 3},
		{airdropId: 2, expectedLength: 2},
	}
	mdl := models.NewModels(gormDb)
	for _, test := range testTable {
		output, err := mdl.GetAllAirdropClaims(test.airdropId)
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
	expected  []models.AirdropClaim
}

func TestGetAllAirdropClaimsReturnsCorrectValues(t *testing.T) {
	prepareTestDatabase()
	tests := []getAllAirdropClaimsTest{
		{1, []models.AirdropClaim{
			{WalletAddress: "0x0", ClaimAmount: "100000", AirdropID: 1},
			{WalletAddress: "0x1", ClaimAmount: "200000", AirdropID: 1},
			{WalletAddress: "0x2", ClaimAmount: "200000", AirdropID: 1},
		}},
		{2, []models.AirdropClaim{
			{WalletAddress: "0x1", ClaimAmount: "200000", AirdropID: 2},
			{WalletAddress: "0x2", ClaimAmount: "400000", AirdropID: 2},
		}},
	}
	mdl := models.NewModels(gormDb)
	for _, test := range tests {
		allClaims, err := mdl.GetAllAirdropClaims(test.airdropId)
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
	expected     models.Airdrop
}

func TestGetLatestAirdropByTokenId(t *testing.T) {
	prepareTestDatabase()
	tests := []getLatestAirdropTest{
		{2, 1, models.Airdrop{ID: 1, TokenID: 2, OnChainIndex: 1, MerkleRoot: "0xRoot"}},
		{3, 2, models.Airdrop{ID: 4, TokenID: 3, OnChainIndex: 2, MerkleRoot: "0xRoot34"}},
		{5, 10, models.Airdrop{ID: 3, TokenID: 5, OnChainIndex: 10, MerkleRoot: "0xRoot4"}},
	}
	mdl := models.NewModels(gormDb)
	for _, test := range tests {
		airdrop, err := mdl.GetAirdropByTokenIndexAndOnChainId(test.tokenId, test.onChainIndex)
		if err != nil {
			t.Error(err)
		}
		if airdrop != test.expected {
			t.Errorf("\nGetAirdropByTokenIndexAndOnChainId() returns correct values for token ID %v, onChainIndex %v: returned %v, want %v", test.tokenId, test.onChainIndex, airdrop, test.expected)
		}
	}
}
