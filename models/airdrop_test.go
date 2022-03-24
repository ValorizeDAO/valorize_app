package models

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/jinzhu/gorm"
)
var (
	db *sql.DB
	fixtures *testfixtures.Loader
	gormDb *gorm.DB
)
func TestMain(m *testing.M) {
	var err error
	dataSourceName := "root:@tcp(127.0.0.1:3306)/valorize_test?charset=utf8&parseTime=True&loc=Local"
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	gormDb, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	gormDb.AutoMigrate(&Airdrop{}, &AirdropClaim{})
	fixtures, err = testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("mysql"),
		testfixtures.Files("fixtures/airdrops.yml"),
		testfixtures.Files("fixtures/airdrop_claims.yml"),
	)
	if err != nil {
		panic(err.Error())
	}
	os.Exit(m.Run())
}
func prepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		panic(err.Error())
	}
}

type getAllAirdropClaimsTestIsRightLength struct {
	airdropId, expectedLength int
}
func TestGetAllAirdropClaimsIsRightLength(t *testing.T) {
	prepareTestDatabase()
	tests := []getAllAirdropClaimsTestIsRightLength{
		{airdropId: 1, expectedLength: 3},
		{airdropId: 2, expectedLength: 2},
	}
	for _, test := range tests{
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
	expected []AirdropClaim
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
	for _, test := range tests{
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