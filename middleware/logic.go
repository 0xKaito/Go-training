package logic

import (
	"example/re/store"
	"log"
	"os"
	"fmt"

	"github.com/joho/godotenv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
)

var data = []byte("unregister")

func Start() (*store.Store, *bind.TransactOpts) {
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

	return instance, auth
}

func CheckSignature(signature []byte) common.Address {
	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex())

	// Recover the public key
	pubKey, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		log.Fatalf("Failed to recover public key: %v", err)
	}

	// Get the signer's address
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	fmt.Printf("Recovered address: %s\n", recoveredAddr.Hex())

	return recoveredAddr;
}