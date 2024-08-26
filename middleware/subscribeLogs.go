package processor

import (
	"example/re/database"
	"example/re/store"
	"example/re/types"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	coreTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"context"
	"strings"

	"github.com/ethereum/go-ethereum"
)


func SubscribLogs(address common.Address, client *ethclient.Client) {
	db := database.Start();
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Error fetching sql db:", err)
        return
    }
	
	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
    }
	
	logs := make(chan coreTypes.Log)
    sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
    if err != nil {
		log.Fatal(err)
    }
	
	contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	if err != nil {
		log.Fatal(err)
	}

	defer sqlDB.Close()

    for {
        select {
        case err := <-sub.Err():
            log.Fatal(err)
        case vLog := <-logs:
			var registerEvent types.LogIsRegistered;
			er := contractAbi.UnpackIntoInterface(&registerEvent, "IsRegistered", vLog.Data);
			if er != nil {
				log.Fatal(er);
			}

			if registerEvent.IsRegistered {
				database.AddUser(db, registerEvent.NewUser);
			} else {
				database.RemoveUser(db, registerEvent.NewUser);
			}
        }
    }
}