package types

import (
	"example/re/store"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)


var SignatureData = []byte("unregister");

type Registration struct {
	Id           int    `gorm:"primaryKey" json:"id"`
	NewUser   string `json:"new_user"`
}

type LogIsRegistered struct {
	NewUser common.Address
    IsRegistered bool;
}

type Config struct {
	ContractAddress string
	ChainId int64
	Client *ethclient.Client
	Instance *store.Store
	Auth *bind.TransactOpts
}