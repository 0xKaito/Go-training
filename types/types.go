package types

import (
	"github.com/ethereum/go-ethereum/common"
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
