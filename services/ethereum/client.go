package ethereum

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"valorize-app/contracts"
	"valorize-app/models"
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

func NewKeystore(password string) (accounts.Account, error) {
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)

	return account, err
}
func _check(e error) {
	if e != nil {
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
		return "", errors.New("could not store keystore on db instance")
	}
	os.Remove(account.URL.Path) // No need to keep the file once in the db.
	return account.Address.Hex(), nil
}

func LaunchContract(client *ethclient.Client, name string, ticker string) (common.Address, *types.Transaction, *contracts.CreatorToken, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(5777))
	auth.Nonce = big.NewInt(int64(000))
	auth.Value = big.NewInt(0)              // in wei
	auth.GasLimit = uint64(300000)          // in units
	auth.GasPrice = big.NewInt(30000000000) // 30 wei

	address, tx, instance, err := contracts.DeployCreatorToken(auth, client, big.NewInt(1000), name, ticker)
	if err != nil {
		return common.HexToAddress("0x0"), nil, nil, err
	}
	return address, tx, instance, nil
}
