package ethereum

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"valorize-app/contracts"
	"valorize-app/models"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jinzhu/gorm"
)

func Connect() (*ethclient.Client, error) {
	client, err := ethclient.Dial(os.Getenv("ETH_NODE"))

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("=======================\n\nConnected to ethereum\n\n=======================")
	}
	return client, err
}

func MainnetConnection() (*ethclient.Client, error) {
	var clientUrl string
	if os.Getenv("ENVIRONMENT") == "PRODUCTION" {
		clientUrl = os.Getenv("MAINNET_NODE")
	} else {
		clientUrl = os.Getenv("ETH_NODE")
	}
	client, err := ethclient.Dial(clientUrl)

	if err != nil {
		log.Println(err.Error())
	} else {
		fmt.Println("=======================\n\nConnected to ethereum mainnet\n\n=======================")
	}
	return client, err
}

func NewKeystore(password string) (accounts.Account, error) {
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)

	return account, err
}

func _check(e error) {
	if e != nil {
		fmt.Println("Error", e.Error())
		panic(e)
	}
}

func StoreUserKeystore(password string, userId uint, DB *gorm.DB) (string, error) {
	account, err := NewKeystore(password)
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

	if DB.Create(&wallet).Error != nil {
		fmt.Println(DB.Create(&wallet).Error.Error())
		return "", errors.New("could not store key")
	}
	os.Remove(account.URL.Path) // No need to keep the file once in the db.
	return account.Address.Hex(), nil
}

func LaunchContract(client *ethclient.Client, name string, ticker string) (common.Address, *types.Transaction, *contracts.CreatorToken, error) {
	fmt.Printf("Launching contract %v (%v)\n\n", name, ticker)
	hotWalletPass := os.Getenv("HOTWALLET_SECRET")
	hotWalletBlob := []byte(os.Getenv("HOTWALLET_KEYSTORE"))
	hotWallet, err := keystore.DecryptKey(hotWalletBlob, hotWalletPass)

	_check(err)
	gasPrice, err := GetGasPrice()
	_check(err)
	txOptions, _ := bind.NewKeyedTransactorWithChainID(hotWallet.PrivateKey, big.NewInt(0003))
	txOptions.Value = big.NewInt(0)      // in wei
	txOptions.GasLimit = uint64(8000000) // in units
	txOptions.GasPrice = big.NewInt(gasPrice)
	fmt.Printf("Gas Price: %v", gasPrice)
	n := new(big.Int)
	initialAmount, _ := n.SetString("1000000000000000000000", 10)
	address, tx, token, err := contracts.DeployCreatorToken(txOptions, client, initialAmount, name, ticker)
	if err != nil {
		fmt.Printf("\nError: %v\n", err.Error())
		return common.HexToAddress("0x0"), nil, nil, err
	}
	return address, tx, token, nil
}

func GetGasPrice() (int64, error) {
	url := os.Getenv("MAINNET_NODE")
	method := "POST"

	payload := strings.NewReader(`{"jsonrpc":"2.0","method":"eth_gasPrice","params": [],"id":1}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	var gasPriceStruct GasPriceStruct
	err = json.Unmarshal(body, &gasPriceStruct)
	_check(err)
	output, err := strconv.ParseInt(hexaNumberToInteger(gasPriceStruct.Result), 16, 64)
	_check(err)
	return output * 5, nil
}

func hexaNumberToInteger(hexaString string) string {
	numberStr := strings.Replace(hexaString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)
	return numberStr
}

type GasPriceStruct struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      uint   `json:"id"`
	Result  string `json:"result"`
}
