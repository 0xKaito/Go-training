package contract

import (
	"example/re/store"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type LogIsRegistered struct {
	NewUser common.Address
    IsRegistered bool;
}

func RegisterUser(instance *store.Store, auth *bind.TransactOpts, address common.Address) {
	result, err := instance.IsRegistered(nil, address)
	if err != nil {
		log.Fatal(err)
	}

	if result {
		log.Fatal("User already registered")
	}

	tx, err := instance.AddUser(auth, address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}

func UnRegister(instance *store.Store, auth *bind.TransactOpts, address common.Address) {
	result, err := instance.IsRegistered(nil, address)
	if err != nil {
		log.Fatal(err)
	}

	if !result {
		log.Fatal("User not registered")
	}

	tx, err := instance.RemoveUser(auth, address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}

func RemoveUser(instance *store.Store, auth *bind.TransactOpts, recoveredAddr common.Address) {

	result, err := instance.IsRegistered(nil, recoveredAddr)
	if err != nil {
		log.Fatal(err)
	}

	if !result {
		log.Fatal("User not registered")
	}

	tx, err := instance.RemoveUser(auth, recoveredAddr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}