package ethereum

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
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
