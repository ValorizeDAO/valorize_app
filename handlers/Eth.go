package handlers

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"os"

	"net/http"
	"valorize-app/models"
	"valorize-app/services"
	"valorize-app/services/ethereum"
)

type EthHandler struct {
	Connection *ethclient.Client
	DB         *gorm.DB
}

func (eth *EthHandler) Ping(c echo.Context) error {
	connection, err := eth.Connection.NetworkID(context.Background())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Could not connect to Ethereum Blockchain",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"data": "connected to " + connection.String(),
	})
}

func (eth *EthHandler) CreateWalletFromRequest(c echo.Context) error {
	password := c.FormValue("password")
	if password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "password required to generate wallet",
		})
	}
	user, _ := services.AuthUser(c, *eth.DB)
	go eth.StoreUserKeystore(password, user.ID)

	return c.JSON(http.StatusOK, map[string]string{
		"status": "wallet is being generated",
	})
}
func _check(e error) {
	if e != nil {
		panic(e)
	}
}

func (eth *EthHandler) StoreUserKeystore(password string, userId uint) (string, error) {
	account, err := ethereum.NewKeystore(password)
	dat, err := ioutil.ReadFile(account.URL.Path)
	_check(err)
	wallet := models.Wallet{
		Address: account.Address.Hex(),
		UserId:  userId,
		Raw:     string(dat),
	}

	if err != nil {
		return "", errors.New("could not create keystore")
	}

	if eth.DB.Create(&wallet).Error != nil {
		return "", errors.New("could not store keystore on db instance")
	}
	os.Remove(account.URL.Path) // No need to keep the file once in the db.
	return account.Address.Hex(), nil
}
