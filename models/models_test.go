package models_test

import (
	"database/sql"
	"os"
	"testing"
	"valorize-app/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/jinzhu/gorm"
)

var (
	db       *sql.DB
	gormDb   *gorm.DB
	fixtures *testfixtures.Loader
)

func TestMain(m *testing.M) {
	db, gormDb = getTestDatabase()
	gormDb.AutoMigrate(&models.Airdrop{}, &models.AirdropClaim{}, models.SmartContract{})
	var err error
	fixtures, err = testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("mysql"),
		testfixtures.Files("fixtures/airdrops.yml"),
		testfixtures.Files("fixtures/airdrop_claims.yml"),
		testfixtures.Files("fixtures/smart_contracts.yml"),
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

func getTestDatabase() (*sql.DB, *gorm.DB) {
	var err error
	dataSourceName := "root:@tcp(127.0.0.1:3306)/valorize_test?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	gormDb, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	return db, gormDb
}
