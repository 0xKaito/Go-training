package main

import (
	processor "example/re/middleware"
	"example/re/router"
	"example/re/types"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var config types.Config;
	config.Instance, config.Auth, config.Client, config.ContractAddress = processor.Start();
	chainId := os.Getenv("CHAIN_ID")
	fmt.Println("sf",chainId);
	intValue, err := strconv.ParseInt(chainId, 10, 64)
    if err != nil {
        log.Fatal("error converting string to int: ", err)
    }
	config.ChainId = intValue;
	router.Start(config);
}
