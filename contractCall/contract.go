package contract

import (
	"example/re/store"
	"fmt"
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func RegisterUser(instance *store.Store, auth *bind.TransactOpts, address common.Address) error {
	result, err := instance.IsRegistered(nil, address)
	if err != nil {
		return err;
	}

	if result {
		return errors.New("user already registered")
	}

	tx, err := instance.AddUser(auth, address)
	if err != nil {
		return err;
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())

	return nil;
}

func UnRegister(instance *store.Store, auth *bind.TransactOpts, address common.Address) error {
	result, err := instance.IsRegistered(nil, address)
	if err != nil {
		return err;
	}

	if !result {
		return errors.New("user not registered")
	}

	tx, err := instance.RemoveUser(auth, address)
	if err != nil {
		return err;
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())

	return nil;
}

func RemoveUser(instance *store.Store, auth *bind.TransactOpts, recoveredAddr common.Address) error {

	result, err := instance.IsRegistered(nil, recoveredAddr)
	if err != nil {
		return err;
	}

	if !result {
		return errors.New("user not registered")
	}

	tx, err := instance.RemoveUser(auth, recoveredAddr)
	if err != nil {
		return err;
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex());

	return nil;
}