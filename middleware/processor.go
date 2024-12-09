package processor

import (
	"example/re/store"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
)

func Start() (*store.Store, *bind.TransactOpts, *ethclient.Client, string) {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	apiURL := os.Getenv("API_URL")
	privateKEY := os.Getenv("PRIVATE_KEY")
	contractAddress := os.Getenv("CONTRACT_ADDRESS")

	client, err := ethclient.Dial(apiURL)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(privateKEY)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(contractAddress)
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	go SubscribLogs(address, client);

	return instance, auth, client, contractAddress
}
