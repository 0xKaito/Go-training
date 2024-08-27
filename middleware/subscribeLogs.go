package processor

import (
	"example/re/database"
	"example/re/store"
	"example/re/types"
	"log"
	"math/big"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	coreTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"context"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/crypto"
)


func SubscribLogs(address common.Address, client *ethclient.Client) (err error) {
	db := database.Start();
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
        return
    }
	
	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
    }
	
	logs := make(chan coreTypes.Log)
    sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
    if err != nil {
		return err
    }
	
	contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	if err != nil {
		return err;
	}

	defer sqlDB.Close()

    for {
        select {
        case err := <-sub.Err():
            return err;
        case vLog := <-logs:
			var registerEvent types.LogIsRegistered;
			er := contractAbi.UnpackIntoInterface(&registerEvent, "IsRegistered", vLog.Data);
			if er != nil {
				return err;
			}

			if registerEvent.IsRegistered {
				err := database.AddUser(db, registerEvent.NewUser);
				if err != nil {
					return err;
				}
			} else {
				err := database.RemoveUser(db, registerEvent.NewUser);
				if err != nil {
					return err;
				}
			}
        }
    }
}

func CheckLogs(contractAddress common.Address, client *ethclient.Client) error {
	query := ethereum.FilterQuery{
        FromBlock: big.NewInt(2394201),
        ToBlock:   big.NewInt(2394201),
        Addresses: []common.Address{
            contractAddress,
        },
    }

    logs, err := client.FilterLogs(context.Background(), query)
    if err != nil {
        log.Fatal(err)
    }

    contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
    if err != nil {
        log.Fatal(err)
    }

    for _, vLog := range logs {
        fmt.Println(vLog.BlockHash.Hex())
        fmt.Println(vLog.BlockNumber) 
        fmt.Println(vLog.TxHash.Hex())

        var registerEvent types.LogIsRegistered;
		er := contractAbi.UnpackIntoInterface(&registerEvent, "IsRegistered", vLog.Data);
		if er != nil {
			return err;
		}

		fmt.Println("Event data")
		fmt.Println(registerEvent.IsRegistered);
		fmt.Println(registerEvent.NewUser);
    }

    eventSignature := []byte("IsRegistered(address,bool)")
    hash := crypto.Keccak256Hash(eventSignature)
    fmt.Println(hash.Hex())

	return nil;
}