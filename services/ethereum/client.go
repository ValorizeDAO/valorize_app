package ethereum

import (
	"context"
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
	creatortoken "valorize-app/contracts_creatortoken"
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
	client, err := ethclient.Dial(os.Getenv("TESTNET_NODE"))

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("=======================\n\nConnected to " + os.Getenv("TESTNET") + "\n\n=======================")
	}
	return client, err
}

func ConnectToChain(chainId string) (*ethclient.Client, error) {
	infuraClientKey := map[string]string{
		"1":      "https://mainnet.infura.io/v3/",
		"3":      "https://ropsten.infura.io/v3/",
		"5":      "https://goerli.infura.io/v3/",
		"137":    "https://polygon-mainnet.infura.io/v3/",
		"10":     "https://optimism-mainnet.infura.io/v3/",
		"421611": "https://arbitrum-rinkeby.infura.io/v3/",
		"42161":  "https://arbitrum-mainnet.infura.io/v3/",
	}
	client, err := ethclient.Dial(infuraClientKey[chainId] + os.Getenv("INFURA_KEY"))
	if err != nil {
		log.Println(err.Error())
	} else {
		fmt.Printf("======\n\n Connected to %s\n\n======\n\n", infuraClientKey[chainId])
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

func LaunchContract(client *ethclient.Client, name string, ticker string, network string) (common.Address, *types.Transaction, *creatortoken.CreatorToken, error) {
	fmt.Printf("Launching contract %v (%v)\n\n", name, ticker)
	hotWalletPass := os.Getenv("HOTWALLET_SECRET")
	hotWalletBlob := []byte(os.Getenv("HOTWALLET_KEYSTORE"))
	hotWallet, err := keystore.DecryptKey(hotWalletBlob, hotWalletPass)

	fmt.Println(network+"_ID:", os.Getenv(network+"_ID"))
	_check(err)
	gasPrice, err := GetGasPrice(network)
	_check(err)
	CHAIN_ID, err := strconv.ParseInt(os.Getenv(network+"_ID"), 10, 64)
	txOptions, _ := bind.NewKeyedTransactorWithChainID(hotWallet.PrivateKey, big.NewInt(CHAIN_ID))
	txOptions.Value = big.NewInt(0)      // in wei
	txOptions.GasLimit = uint64(3000000) // in units
	txOptions.GasPrice = big.NewInt(gasPrice)
	fmt.Printf("Gas Price: %v\n", gasPrice)
	n := new(big.Int)
	initialAmount, _ := n.SetString("1000000000000000000000", 10)

	//see if balance is enough to launch
	balance, err := client.BalanceAt(context.Background(), hotWallet.Address, nil)
	fmt.Println("Balance: ", balance)
	gasCost := big.NewInt(0).Mul(txOptions.GasPrice, big.NewInt(3000000))
	fmt.Println("Gas Cost: ", gasCost)
	if balance.Cmp(gasCost) == -1 {
		return common.Address{}, nil, nil, errors.New("Not enough balance to launch contract")
	}

	address, tx, token, err := creatortoken.DeployCreatorToken(txOptions, client, initialAmount, name, ticker, hotWallet.Address)
	if err != nil {
		fmt.Printf("\nError: %v\n", err.Error())
		return common.HexToAddress("0x0"), nil, nil, err
	}
	return address, tx, token, nil
}

func GetGasPrice(network string) (int64, error) {
	var url string
	if network == "MAINNET" {
		url = os.Getenv("MAINNET_NODE")
	} else if network == "TESTNET" {
		url = os.Getenv("TESTNET_NODE")
	} else {
		return 0, errors.New("invalid network")
	}

	payload := strings.NewReader(`{"jsonrpc":"2.0","method":"eth_gasPrice","params": [],"id":` + os.Getenv(network+"_ID") + `}`)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

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
